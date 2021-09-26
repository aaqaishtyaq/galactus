package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	DEFAULT_RELEASE_TYPE = "patch"
)

type releaseTag struct {
	gitSha  string
	service string
	version string
}

type Version struct {
	major int
	minor int
	patch int
}

func main() {
	if len(os.Args) > 1 {
		srvArgs := os.Args[1:]
		srvStock := getServiceNames()
		srvToBuild := make([]string, 0)

		if len(srvArgs) == 1 && srvArgs[0] == "all" {
			srvToBuild = srvStock
		} else {
			for _, srv := range srvArgs {
				for _, tartgetSrv := range srvStock {
					if srv == tartgetSrv {
						srvToBuild = append(srvToBuild, srv)
					}
				}
			}
		}

		if err := buildImages(srvToBuild); err != nil {
			log.Fatal(err)
		}

	} else {
		fmt.Println("Please Provide required args")
		return
	}
}

func getServiceNames() []string {
	basePath, err := getRepoBasePath()
	if err != nil {
		log.Fatal(err)
	}

	absBasePath := strings.TrimRight(basePath, "/hack/")
	files, err := ioutil.ReadDir(absBasePath + "/components")
	if err != nil {
		log.Fatal(err)
	}

	serviceNames := make([]string, len(files))

	for i, file := range files {
		serviceNames[i] = file.Name()
	}

	return serviceNames
}

func getRepoBasePath() (string, error) {
	basePath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return basePath, nil
}

func buildImages(images []string) error {
	for _, img := range images {
		newRel := serviceTag(img)
		releaseService(img, newRel)
	}
	return nil
}

func serviceTag(image string) string {
	currentTag := getCurrentGitTag(image)
	var releaseTagInfo *releaseTag
	if len(currentTag) < 1 {
		releaseTagInfo = &releaseTag{
			gitSha:  "",
			service: image,
			version: "v0.0.1",
		}
	} else {
		releaseTagInfo = extractRelease(currentTag)
		bumpRelease(releaseTagInfo)
	}

	setGitSha(releaseTagInfo)
	return fmt.Sprintf("release#%v#%v#%v", releaseTagInfo.service, releaseTagInfo.version, releaseTagInfo.gitSha)
}

func releaseService(serviceImage, tag string) {
	fmt.Println(tag)
	stdout, _, err := runShell(
		"git",
		"tag",
		"-a",
		tag,
		"-m",
		"\"Release for SERVICE \"",
	)

	if err != nil {
		fmt.Println("Unable to release image tag")
		log.Fatal(err)
	}

	fmt.Println(stdout.String())
}

func getCurrentGitTag(service string) string {
	stdout, _, err := runShell(
		"git",
		"tag",
		"-l",
		"*"+service+"*",
	)
	if err != nil {
		log.Fatal(err)
	}

	// trim whitespaces from the stdout
	// split it based on new line
	// and select the last tag release
	tags := strings.Split(strings.TrimSpace(stdout.String()), "\n")
	return tags[len(tags)-1]
}

func extractRelease(gitTag string) *releaseTag {
	tagInfo := strings.Split(gitTag, "#")
	gitTagInfo := &releaseTag{
		service: tagInfo[1],
		version: tagInfo[2],
		gitSha:  tagInfo[3],
	}

	return gitTagInfo
}

func bumpRelease(tag *releaseTag) {
	currentVersion := strings.TrimLeft(tag.version, "v")
	version := bumpVersion(currentVersion, DEFAULT_RELEASE_TYPE)
	newVersion := fmt.Sprintf("v%v.%v.%v", version.major, version.minor, version.patch)
	tag.version = newVersion
}

func bumpVersion(current, releaseType string) *Version {
	deconsVer := strings.Split(current, ".")
	deconsVerInt := make([]int, 3)
	for i, v := range deconsVer {
		vInt, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		deconsVerInt[i] = vInt
	}

	version := &Version{
		major: deconsVerInt[0],
		minor: deconsVerInt[1],
		patch: deconsVerInt[2],
	}

	switch releaseType {
	case "major":
		version.major = (version.major + 1)
	case "minor":
		version.minor = (version.minor + 1)
	case "patch":
		version.patch = (version.patch + 1)
	default:
		version.patch = (version.patch + 1)
	}

	return version
}

func setGitSha(rTag *releaseTag) *releaseTag {
	stdout, _, err := runShell(
		"git",
		"rev-parse",
		"--short",
		"HEAD",
	)
	if err != nil {
		log.Fatal(err)
	}

	rTag.gitSha = strings.TrimSpace(stdout.String())

	return rTag
}

func runShell(cmd string, args ...string) (bytes.Buffer, bytes.Buffer, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	shell := exec.Command(cmd, args...)
	shell.Stdout = &stdout
	shell.Stderr = &stderr

	if err := shell.Run(); err != nil {
		return stdout, stderr, err
	}

	return stdout, stderr, nil
}

const form = document.querySelector("form");
const input = document.querySelector("input");
const span = document.querySelector("span");

const socket = new WebSocket("ws://localhost:7000/workspace");

socket.addEventListener("open", () => {
  span.innerHTML = span.innerHTML + `<br/> Websocket connection initialised with ${socket.url}`
});

socket.addEventListener("close", () => {
  span.innerHTML = span.innerHTML + `<br/> Websocket connection terminated from ${socket.url}`
});

socket.addEventListener("message", (e) => {
  span.innerHTML = span.innerHTML + `<br/> ${e.data}`
});

form.addEventListener("submit", (e) => {
  e.preventDefault()
  sendMessage(input.value)
  input.value = ""
})

function sendMessage(message) {
  span.innerHTML =
    span.innerHTML + `<br/>Workspace request sent for the following Git Remote URL: "${message}"`
  socket.send(message);
}

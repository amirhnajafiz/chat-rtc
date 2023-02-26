const input = document.querySelector('#textarea')
const messages = document.querySelector('#messages')
const username = document.querySelector('#username')
const send = document.querySelector('#send')

const url = "ws://" + window.location.host + "/ws";
const ws = new WebSocket(url);

ws.onmessage = function (msg) {
    insertMessage(JSON.parse(msg.data))
};

send.onclick = () => {
    const message = {
        username: username.value,
        content: input.value,
    }

    ws.send(JSON.stringify(message));
    input.value = "";
};

function insertMessage(messageObj) {
    const message = document.createElement('div')

    message.setAttribute('class', 'chat-message')
    console.log("name: " +messageObj.username + " content: " + messageObj.content)
    message.textContent = `${messageObj.username}: ${messageObj.content}`

    messages.appendChild(message)

    messages.insertBefore(message, messages.firstChild)
}
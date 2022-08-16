<h1 align="start">
  Robin.io-go
</h1>


## Table of contents

<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#introduction">Introduction</a>
    </li>
    <li>
      <a href="#prerequisites">Prerequisites</a>
    </li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#sending-your-first-message">Sending your first message</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>

<br />

## Introduction

Robin.io-go is The official Go SDK built to communicate with the [Robinapp API](https://robinapp.co/). Now you can integrate [Robin.io](https://robinapp.co/) with minimal effort and quickly setup a real-time messaging platform in your Web application.

## Prerequisites

The following packages are required to use the sdk:
* go 1.16
* Robin api_key

## Getting started

#### Step 1: Create a Robinapp account

A Robinapp account comprises everything required in a chat service including users, message, and api-keys. To create an application:

1. Go to the [Robinapp Dashboard](https://dashboard.robinapp.co/signup) and enter your email and password, and create a new account.
2. Navigate to [Api Config](https://dashboard.robinapp.co/apiconfig) and copy your `API key`

> Note: All the data is limited to the scope of a single user account, thus the users in different Robinapp accounts are unable to chat with each other.

#### Step 2: Install the SDK

```
  go get robin.io-go
```

## Sending your first message

Follow the step-by-step instructions below to authenticate and send your first message.

### Authentication

To use the features of the Chat SDK in your client app, a `robin` instance must be initiated in each client app before user authentication with Robin server. These instances communicate and interact with the server based on an authenticated user account, allowing for the client app to use the Chat SDK features.

### Step 1: Initialize the Chat SDK

To initialize a `Robin` instance, pass the `API key` as the first argument to in the `Robin()` method, You can find your API key in the API Configuration tab in your [Robin Account](https://robin-user.herokuapp.com/apiconfig).

Then `true` or `false` for as the second parameter as it tells the sdk whether to load with ssl or not.

```golang
robin := Robin{
		Secret: "NT-qBsdC.....kOBVYRr",
		Tls:    true,
	}
```

### Step 2: Connect to Robin server

You'll need a **USER_TOKEN** to connect to the Robin server.

#### A. Create User Token

Create user token

```go
...
token, err := robin.CreateUserToken(UserToken{MetaData: map[string]interface{}{
		"name": "michael",
	}})

if err != nil {
    println(err)
}
```

Connect to the Robin server using the **USER_TOKEN** you just created.

```go
// call back for a successful connection
func connected(socket gowebsocket.Socket) {}
// call back for an unsuccessful connection
func disconnected(err error, socket gowebsocket.Socket) {}
// call back for when a message is recieved via the connection
func text_recieved(msg string, socket gowebsocket.Socket) {}
conn, err := robin.Connect(connected, nil, disconnected, text_recieved, nil, nil, nil)

if err != nil {
    println(err)
}
```

### Step 3: Channels

All messages sent via Robin are sent through channels, you can consider channels as tunnels that relay messages to all connected clients.

### Step 4: Create a conversation

Before we can send a message to a channel we first need to create a converstion.

```go
robin := Robin{
    Secret: "NT-QuNtKo......JwGrymaVxQX",
    Tls:    true,
}
/* 
func (*Robin).CreateConversation(senderName string, senderToken string, receiverToken string, receiverName string) (ConversationResponseData, error)
*/
conv, err := robin.CreateConversation("elvis", "YFXOK....BaqKgDWOhE", "YFXOKV...gDWOhE", "jesse")

if err != nil {
    fmt.Println(err)
}

// create group conversation

notify := Robin{
    Secret: "NT-QuNtKolp.....cJwGrymaVxQX",
    Tls:    true,
}

/* 
func (*Robin).CreateGroupConversation(name string, moderator UserToken, participants []UserToken) (ConversationResponseData, error)
*/

conv, err := notify.CreateGroupConversation("The council",
    UserToken{UserToken: "YFXOKVyKBGv...KgDWOhE"},
    []UserToken{{UserToken: "TKSSAK...gDWOhE"},
    })

if err != nil {
    t.Error(err)
}
```

### Step 5: Subscribe to a channel

To send and recieve messages in Robin, we utilize channels, your connection has to be subscribed to a channel to send a recieve messages through and from it.

```go
err := robin.Subscribe("<channel_name>")
if err != nil {
    fmt.Println(err)
}
```

### Step 6: Send a message

Finally, send a message

```go
err := robin.SendMessage("<channel_name>", map[string]interface{}{
    "name":"elvis",
    "user_token":"TKSHSqww....aA",
    "msg":"hello from robin",
})
```

#### Options

The following are general attributes used in Robin:

|   Attribute    |   Type    | Default | Description                                              |
| :------------: | :-------: | :-----: | :------------------------------------------------------- |
|      socket      | gowebsocket.Socket |  nil   | Websocket object returned after calling robin.connect() |
|  sender_name   |  String   |   ''    | Name of the person sending the message                   |
|  sender_token  |  String   |   ''    | USER_TOKEN of the person sending the message             |
| receiver_name  |  String   |   ''    | Name of the person receiving the message                 |
| receiver_token |  String   |   ''    | USER_TOKEN of the person receiving the message           |
|      msg       |  map[string]interface{}   |   {}    | Json serializable object containing the message          |

If you have any comments or questions regarding bugs and feature requests, visit [Robinapp community](https://community.robinapp.co).

[View Documentation Here]().

## License

Distributed under the MIT License. See `LICENSE` for more information.

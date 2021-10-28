# message-sender
This telegram bot can send messages with selected priority

It is a server that sends messages to a telegram channel/group using a telegram bot token.

Steps that should be taken to run server:

1) create .env file inside which there should be:
{
	HTTP_PORT = :8000

	GRPC_SERVER_PORT = :8008


	TELEGRAM_APITOKEN = 
	CHAT_ID = 
}

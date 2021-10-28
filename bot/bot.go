package bot

import (
    "os"
    "github.com/joho/godotenv"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "log"
    "strconv"
)

func BotService(text string) (err error) {
    err = godotenv.Load()
    if err != nil {
        log.Fatalf("Problem with loading env: %v", err)
        return 
    }

    bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
    if err != nil {
        log.Fatalf("Problem with connecting to bot: %v", err)
        return 
    }

    chatId := os.Getenv("CHAT_ID")
    id, err := strconv.ParseInt(chatId, 10, 64)
    if err != nil {
        log.Fatalf("Problem with converting chatId to int64: %v", err)
        return 
    }

    bot.Debug = true
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
   
    msg := tgbotapi.NewMessage(id,text)
    if _, err := bot.Send(msg); err != nil {
        log.Fatalf("Problem with sending message: %v",err)
        return err
    }
        
    return 
}
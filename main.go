package main

import (
  "net/http"
  "os"
  "log"
  "strconv"

  "github.com/gin-gonic/gin"
)

func SubmitToMailChimp(name string, email string, uType string) int {
  log.Println("name:", name, "email:", email, "uType:", uType)
  req, err := http.NewRequest("POST", os.Getenv("MC_URL") + "/lists/" + os.Getenv("MC_LIST_ID") + "/members", nil)
  if err != nil {
    log.Fatal("NewRequest: ", err)
    return http.StatusInternalServerError
  }

  req.Header.Add("Authorization", "apikey " + os.Getenv("MC_API_KEY"))

  client := &http.Client{}

  resp, err := client.Do(req)
  if err != nil {
    log.Println("Do: ", err)
    return http.StatusInternalServerError
  }

  defer resp.Body.Close()
  log.Println(resp.Body)
  returnStatus := (resp.Status)[:3]
  stat, err := strconv.Atoi(returnStatus)

  if err != nil {
    log.Println("Status parse error")
    return http.StatusInternalServerError
  }

  return stat
}

func main() {
  // Creates a gin router with default middleware:
  // logger and recovery (crash-free) middleware
  router := gin.Default()
  router.Use(gin.Logger())
  router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static") // For static assets

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl.html",nil)
  })

  router.GET("/register_user", func(c *gin.Context) {
    // Make API Call
    status := SubmitToMailChimp(c.Query("name"), c.Query("email"), c.Query("uType"))
    // Respond
    var msg struct {
    			Message string
    			Status  int
    }

    // Added user successfully
    msg.Message = "Student added."
    msg.Status = status
/*
    // User already on list
    msg.Message = "Student already on mailing list."
    msg.Status = http.statusFound

    // Error
    msg.Message = "Unkown Error."
    msg.Status = http.statusBadRequest
*/
    c.JSON(http.StatusOK, msg)
  })
  // By default it serves on :8080 unless a
  // PORT environment variable was defined.
  router.Run()
}

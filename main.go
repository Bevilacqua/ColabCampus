package main

import (
  "net/http"
  "os"
  "log"
  "strconv"
  "bytes"

  "github.com/gin-gonic/gin"
)

func SubmitToMailChimp(name string, email string, uType string) int {
  // Encode Body
  var bodyStr string = `{"email_address":"` + email + `","status": "subscribed","merge_fields": {"NAME":"` + name + `","UTYPE":"` + uType + `"}}`
  var body = []byte(bodyStr)

  // Make request
  req, err := http.NewRequest("POST", os.Getenv("MC_URL") + "/lists/" + os.Getenv("MC_LIST_ID") + "/members", bytes.NewBuffer(body))
  if err != nil {
    log.Fatal("NewRequest: ", err)
    return http.StatusInternalServerError
  }
  log.Println(string(bodyStr))
  req.Header.Add("Authorization", "apikey " + os.Getenv("MC_API_KEY"))

  // Add client
  client := &http.Client{}

  resp, err := client.Do(req)
  if err != nil {
    log.Println("Do: ", err)
    return http.StatusInternalServerError
  }

  defer resp.Body.Close()

  // Check return code
  returnStatus := (resp.Status)[:3]
  stat, err := strconv.Atoi(returnStatus)

  if err != nil {
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
    // Validate Input!
/*
    msg.Message = "Missing Information"
    msg.Status = http.StatusBadRequest
*/

    // Make API Call
    status := SubmitToMailChimp(c.Query("name"), c.Query("email"), c.Query("uType"))
    // Respond
    var msg struct {
    			Message string
    			Status  int
    }

    // Added user successfully
    if status == 200 {
      msg.Message = "Student added."
      msg.Status = status
    } else {
      // User already on list
      msg.Message = "Student already on mailing list."
      msg.Status = http.StatusFound
    }

    c.JSON(http.StatusOK, msg)
  })
  // By default it serves on :8080 unless a
  // PORT environment variable was defined.
  router.Run()
}

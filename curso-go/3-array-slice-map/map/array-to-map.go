package main

import "fmt"

type NotificationInfoVariable struct {
  Name  string `json:"name"`
  Value string `json:"value"`
}

func arrayToMap(notifications []NotificationInfoVariable) map[string]string {
  notificationMap := make(map[string]string)
  for _, notification := range notifications {
    notificationMap[notification.Name] = notification.Value
  }
  return notificationMap
}

func main() {
  // Sample array of NotificationInfoVariable structs
  notifications := []NotificationInfoVariable{
    {Name: "Status", Value: "Active"},
    {Name: "Level", Value: "Warning"},
  }
  
  notificationMap := arrayToMap(notifications)
	fmt.Println(notificationMap["Level"])
  fmt.Println(notificationMap) // Output: map[Level:Warning Status:Active]
}

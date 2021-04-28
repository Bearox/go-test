package main

import "fmt"

func main() {
	topicEventMap := make(map[string][]string, 1)

	monitorLogTopic := "monitor-log"

	addEvent(monitorLogTopic, topicEventMap, "hello")
	addEvent(monitorLogTopic, topicEventMap, "bearox")

	hunterTopic := "hunter-log"
	addEvent(hunterTopic, topicEventMap, "hi")
	addEvent(hunterTopic, topicEventMap, "weiwen")

	fmt.Println(topicEventMap)

	for _, events:= range topicEventMap {
		fmt.Println(events)
	}
}

func addEvent(topic string, topicEventMap map[string][]string, event string)  {
	events, ok := topicEventMap[topic]

	if ok {
		events = append(events, event)
	} else {
		events = []string{event}
		topicEventMap[topic] = events
	}
	topicEventMap[topic] = events
}

package main

func main() {
	ok := cmd.Init()
	if ok == false {
		return
	}
	producer.Init()

	go producer.Start()
	go consumer.Start()

	task.Wait()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

type User interface {
	PrintName()
	PrintDetails()
}

type Person struct {
	FirstName string
	LastName  string
	Dob       time.Time
	Email     string
	Location  string
}

//A person method
func (p Person) PrintName(newName string) {
	p.FirstName = newName
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

//A person method
func (p Person) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s ]\n", p.Dob.String(), p.Email,
		p.Location)
}

//A person method with pointer receiver
func (p *Person) ChangeName(newName string) {
	p.FirstName = newName
}

//A person method with pointer receiver
func (p *Person) ChangeLocation(newLocation string) {
	p.Location = newLocation
}

type Admin struct {
	Person //type embedding for composition
	Roles  []string
}

//overrides PrintDetails
func (a Admin) PrintDetails() {
	//Call person PrintDetails
	a.Person.PrintDetails()
	fmt.Println("Admin Roles:")
	for _, v := range a.Roles {
		fmt.Println(v)
	}
}

type Member struct {
	Person //type embedding for composition
	Skills []string
}

//overrides PrintDetails
func (m Member) PrintDetails() {
	//Call person PrintDetails
	m.Person.PrintDetails()
	fmt.Println("Skills:")
	for _, v := range m.Skills {
		fmt.Println(v)
	}
}

type Team struct {
	Name, Description string
	Users             []User
}

func (t Team) GetTeamDetails() {
	fmt.Printf("Team: %s  - %s\n", t.Name, t.Description)
	fmt.Println("Deteails of the team members:")
	for _, v := range t.Users {
		v.PrintName()
		v.PrintDetails()
	}
}

// wg is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func main() {
	var p Person
	p.FirstName = "Shiju"
	p.LastName = "Varghese"
	p.Dob = time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC)
	p.Email = "shiju@email.com"
	p.Location = "Kochi"

	budi := Person{
		FirstName: "Shiju",
		LastName:  "Varghese",
		Dob:       time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
		Email:     "shiju@email.com",
		Location:  "Kochi",
	}

	fmt.Println(p)
	fmt.Println(budi)

	asep := Person{
		"Shiju",
		"Varghese",
		time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
		"shiju@email.com",
		"Kochi",
	}
	fmt.Println(asep)

	tuti := Person{FirstName: "Shiju", LastName: "Varghese"}
	p.Dob = time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC)
	p.Email = "shiju@email.com"
	p.Location = "Kochi"
	fmt.Println(tuti)

	tuti.PrintName("Budi Setiono")
	tuti.PrintDetails()

	fmt.Println("Nama : ", tuti.FirstName)

	cecep := &Person{FirstName: "Shiju", LastName: "Varghese"}
	p.Dob = time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC)
	p.Email = "shiju@email.com"
	p.Location = "Kochi"

	cecep.ChangeName("Cecep")
	fmt.Println(&cecep)

	deden := &Person{

		"Shiju",
		"Varghese",
		time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
		"shiju@email.com",
		"Kochi",
	}
	deden.ChangeLocation("Bandung")
	deden.PrintName("Deden")
	deden.PrintDetails()

	alex := Admin{
		Person{
			"Alex",
			"John",
			time.Date(1970, time.January, 10, 0, 0, 0, 0, time.UTC),
			"alex@email.com",
			"New York"},
		[]string{"Manage Team", "Manage Tasks"},
	}
	shiju := Member{
		Person{
			"Shiju",
			"Varghese",
			time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
			"shiju@email.com",
			"Kochi"}, []string{"Go", "Docker", "Kubernetes"},
	}

	//call methods for alex
	alex.PrintName("Alex")
	alex.PrintDetails()

	//call methods for shiju
	shiju.PrintName("Shiju")
	shiju.PrintDetails()

	// Add a count of two, one for each goroutine.
	// wg.Add(2)
	// fmt.Println("Start Goroutines")

	//launch a goroutine with label "A"
	// go printCounts("A")

	//launch a goroutine with label "B"
	// go printCounts("B")

	// Wait for the goroutines to finish.
	// fmt.Println("Waiting To Finish")
	// wg.Wait()

	// fmt.Println("\nTerminating Program")

	// UnBuffered Channel
	count := make(chan int)

	// Add a count of two, one for each goroutine.
	wg.Add(2)
	fmt.Println("Start Goroutines")

	//launch a goroutine with label "B"
	go printCounts("B", count)

	//launch a goroutine with label "A"
	go printCounts("A", count)

	fmt.Println("Channel begin")
	count <- 1

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")

	messages := make(chan string, 2)
	messages <- "Golang"
	messages <- "Gopher"
	//Recieve value from buffered channel
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	// fmt.Println(<-messages)
	messages <- "Golang"
	messages <- "Gopher"
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}

// End Of Main

// func printCounts(label string) {
// 	// Schedule the call to WaitGroup's Done to tell we are done.
// 	defer wg.Done()

// 	// Randomly wait
// 	for count := 1; count <= 10; count++ {
// 		sleep := rand.Int63n(1000)
// 		time.Sleep(time.Duration(sleep) * time.Millisecond)
// 		fmt.Printf("Count: %d from %s\n", count, label)
// 	}
// }

func printCounts(label string, count chan int) {
	// Schedule the call to WaitGroup's Done to tell we are done.
	defer wg.Done()
	for {
		//Receives message from Channel
		val, ok := <-count
		if !ok {
			fmt.Println("Channel was closed")
			return
		}
		fmt.Printf("Count: %d received from %s \n", val, label)
		if val == 10 {
			fmt.Printf("Channel Closed from %s \n", label)
			// Close the channel
			close(count)
			return
		}
		val++
		// Send count back to the other goroutine.
		count <- val
	}
}

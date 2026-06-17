// package main

// import "fmt"

// func main() {
// 	//fmt.Println("Hello, World!");
	
// 	var conferenceName string = "Go Conference"
// 	const conferenceTickets = 50
// 	var remainingTickets = 50

// 	fmt.Printf("Welcome to %v\n", conferenceName)
// 	fmt.Println("Get your tickets here to attend the conference.");
// 	fmt.Printf("There are %v tickets remaining.\n", remainingTickets)

// 	for {

// 		fmt.Printf("Enter your first name: ")
// 		var firstName string
// 		fmt.Scan(&firstName)
// 		fmt.Printf("Enter the number of tickets you want to book: ")
// 		var numTickets int
// 		fmt.Scan(&numTickets)
// 		fmt.Printf("Thank you %v for booking %v tickets.\n", firstName, numTickets)
		
// 		// Array example
// 		var bookings [50]string
// 		bookings[0] = firstName
// 		fmt.Printf("The first booking is: %v\n", bookings[0])

// 		// Slice example
// 		// var bookingsSlice []string
// 		// bookingsSlice = append(bookingsSlice, firstName)
// 		// fmt.Printf("The first booking in the slice is: %v\n", bookingsSlice[0])
// 	}
	
// }
package main

import "fmt"

// Struct — your first essential: composite data types
type Contact struct {
    ID    int
    Name  string
    Email string
    Phone string
}

// Map as the "database" — key: ID, value: Contact
var contacts = make(map[int]Contact)
var nextID = 1
// CREATE — pass by value, store in map
func addContact(name, email, phone string) {
    c := Contact{ID: nextID, Name: name, Email: email, Phone: phone}
    contacts[nextID] = c
    nextID++
    fmt.Printf("✅ Contact added with ID %d\n", c.ID)
}

// READ ALL — range over map (notice: unordered!)
func listContacts() {
    if len(contacts) == 0 {
        fmt.Println("📭 No contacts found.")
        return
    }
    for _, c := range contacts {
        fmt.Printf("[%d] %s | %s | %s\n", c.ID, c.Name, c.Email, c.Phone)
    }
}

// UPDATE — pointer concept: why we re-assign to map
func updateContact(id int, newPhone string) {
    c, exists := contacts[id]   // comma-ok idiom
    if !exists {
        fmt.Println("❌ Contact not found.")
        return
    }
    c.Phone = newPhone
    contacts[id] = c            // maps store copies — must re-assign!
    fmt.Println("✏️  Contact updated.")
}

// DELETE — built-in delete()
func deleteContact(id int) {
    if _, exists := contacts[id]; !exists {
        fmt.Println("❌ Contact not found.")
        return
    }
    delete(contacts, id)
    fmt.Println("🗑️  Contact deleted.")
}
func showMenu() {
    fmt.Println("\n===== 📒 Contact Book =====")
    fmt.Println("1. Add Contact")
    fmt.Println("2. List Contacts")
    fmt.Println("3. Update Phone")
    fmt.Println("4. Delete Contact")
    fmt.Println("5. Exit")
    fmt.Print("Choose: ")
}

func main() {
    var choice int

    for {                          // infinite loop — Go's while
        showMenu()
        fmt.Scan(&choice)          // & = pointer to choice

        switch choice {
        case 1:
            var name, email, phone string
            fmt.Print("Name: ")  
			fmt.Scan(&name)
            fmt.Print("Email: ") 
			fmt.Scan(&email)
            fmt.Print("Phone: ") 
			fmt.Scan(&phone)
            addContact(name, email, phone)
        case 2:
            listContacts()
        case 3:
            var id int; var phone string
            fmt.Print("ID: ")    
			fmt.Scan(&id)
            fmt.Print("New Phone: ") 
			fmt.Scan(&phone)
            updateContact(id, phone)
        case 4:
            var id int
            fmt.Print("ID to delete: ") 
			fmt.Scan(&id)
            deleteContact(id)
        case 5:
            fmt.Println("👋 Goodbye!")
            return
        default:
            fmt.Println("⚠️  Invalid option.")
        }
    }
}
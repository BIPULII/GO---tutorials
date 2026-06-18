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
// var contacts = make(map[int]Contact)
var contacts []Contact
var nextID = 1
// CREATE — pass by value, store in map
func addContact(name, email, phone string) {
    c := Contact{ID: nextID, Name: name, Email: email, Phone: phone}
    contacts = append(contacts, c)
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
func updateContact(id int, newPhone string , newName string, newEmail string) {
    for i := 0; i < len(contacts); i++ {
        if contacts[i].ID == id {
            contacts[i].Phone = newPhone
            contacts[i].Name = newName
            contacts[i].Email = newEmail
            fmt.Println("✏️  Contact updated.")
            return
        }
    }
    fmt.Println("❌ Contact not found.")
}

// DELETE — built-in delete()
func deleteContact(id int) {
    for i := 0; i < len(contacts); i++ {
        if contacts[i].ID == id {
            contacts = append(contacts[:i], contacts[i+1:]...)
            fmt.Println("🗑️  Contact deleted.")
            return
        }
    }
    fmt.Println("❌ Contact not found.")
}
func searchContact(name string) {
    found := false

    for _, c := range contacts {
        if c.Name == name {
            fmt.Printf("[%d] %s | %s | %s\n",
                c.ID, c.Name, c.Email, c.Phone)
            found = true
        }
    }

    if !found {
        fmt.Println("Contact not found.")
    }
}
func showMenu() {
    fmt.Println("\n===== 📒 Contact Book =====")
    fmt.Println("1. Add Contact")
    fmt.Println("2. List Contacts")
    fmt.Println("3. Update Phone")
    fmt.Println("4. Delete Contact")
    fmt.Println("5. Exit")
	fmt.Println("6. Search Contact")
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
            var id int; var phone string; var name string; var email string
            fmt.Print("ID: ")    
			fmt.Scan(&id)
            fmt.Print("New Phone: ") 
			fmt.Scan(&phone)
            fmt.Print("New Name: ") 
			fmt.Scan(&name)
            fmt.Print("New Email: ") 
			fmt.Scan(&email)
            updateContact(id, phone, name, email)
        case 4:
            var id int
            fmt.Print("ID to delete: ") 
			fmt.Scan(&id)
            deleteContact(id)
        case 5:
            fmt.Println("👋 Goodbye!")
            return
		case 6:
				var name string
				fmt.Print("Search Name: ")
				fmt.Scan(&name)
				searchContact(name)
        default:
            fmt.Println("⚠️  Invalid option.")
        }
    }
}
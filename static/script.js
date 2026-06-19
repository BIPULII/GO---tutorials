let contacts = [];

async function loadContacts() {

    const response =
        await fetch("/contacts");

    contacts =
        await response.json();

    renderTable(contacts);
}

function renderTable(data) {

    const table =
        document.getElementById(
            "contactTable"
        );

    table.innerHTML = "";

    document.getElementById(
        "count"
    ).innerText = data.length;

    data.forEach(c => {

        table.innerHTML += `

        <tr>

            <td>${c.id}</td>

            <td>${c.name}</td>

            <td>${c.email}</td>

            <td>${c.phone}</td>

            <td>

                <button
                class="edit-btn"
                onclick="editContact(${c.id})">
                Edit
                </button>

                <button
                class="delete-btn"
                onclick="deleteContact(${c.id})">
                Delete
                </button>

            </td>

        </tr>

        `;
    });
}

async function addContact() {

    const name =
        document.getElementById(
            "name"
        ).value;

    const email =
        document.getElementById(
            "email"
        ).value;

    const phone =
        document.getElementById(
            "phone"
        ).value;

    if (!name || !email || !phone) {
        alert("Fill all fields");
        return;
    }

    await fetch("/add", {

        method: "POST",

        headers: {
            "Content-Type":
            "application/json"
        },

        body: JSON.stringify({
            name,
            email,
            phone
        })
    });

    document.getElementById("name").value = "";
    document.getElementById("email").value = "";
    document.getElementById("phone").value = "";

    loadContacts();
}

async function deleteContact(id) {

    if (
        !confirm(
            "Delete this contact?"
        )
    ) return;

    await fetch(
        `/delete?id=${id}`
    );

    loadContacts();
}

async function editContact(id) {

    const contact =
        contacts.find(
            c => c.id === id
        );

    const newName =
        prompt(
            "Enter new name",
            contact.name
        );

    if (!newName) return;

    const newEmail =
        prompt(
            "Enter new email",
            contact.email
        );

    const newPhone =
        prompt(
            "Enter new phone",
            contact.phone
        );

    await fetch("/update", {

        method: "POST",

        headers: {
            "Content-Type":
            "application/json"
        },

        body: JSON.stringify({

            id,

            name: newName,

            email: newEmail,

            phone: newPhone
        })
    });

    loadContacts();
}

function searchContact() {

    const value =
        document
        .getElementById(
            "search"
        )
        .value
        .toLowerCase();

    const filtered =
        contacts.filter(c =>
            c.name
            .toLowerCase()
            .includes(value)
        );

    renderTable(filtered);
}

window.onload = loadContacts;
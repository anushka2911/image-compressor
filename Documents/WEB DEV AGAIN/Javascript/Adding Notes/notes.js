let inputDate=document.getElementById("date")
let inputNotes=document.getElementById("notes")

let date;
let notes;
let add = () => {
    let date = inputDate.value;
    let notes = inputNotes.value;

    appendDateAndNotes(date, notes);
    
}

let appendDateAndNotes = (date, notes) => {
    let noteContainer = document.createElement("div"); // Create a container div
    noteContainer.setAttribute("class", "note"); 

    let notePadDate = document.createElement("span");
    notePadDate.setAttribute("id", "dateByUser");
    notePadDate.innerText = date;

    let notePadNotes = document.createElement("p");
    notePadNotes.setAttribute("class", "stickyNotes");
    notePadNotes.innerText = notes;

    noteContainer.appendChild(notePadDate); 
    noteContainer.appendChild(notePadNotes);



    let appendArea=document.getElementById("AppendArea")
    appendArea.append(noteContainer)
}



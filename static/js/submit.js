const createSubmitBtn = document.getElementById('create_submit')
const insertForm = document.getElementById('insert_values')

createSubmitBtn.addEventListener('click', () => {
    const xhr = new XMLHttpRequest()
    xhr.open("POST", "http://localhost:8080/create_submit" + "?drvalue=" + dropDown.value + "&values=" + insertForm.value)
    xhr.send()
    xhr.onload = () => {
        location.reload()
    } 
})
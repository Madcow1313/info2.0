const crudButtons = document.getElementsByClassName('CRUD_buttons')
const crudMap = new Map()



crudMap.set('create_btn', 'create_fields')
crudMap.set('read_btn', 'read_fields')
crudMap.set('delete_btn', 'delete_fields')
crudMap.set('update_btn', 'update_fields')


for (const btn of crudButtons) {
    btn.addEventListener('click', () => {
        for (const el of crudMap) {
			console.log(btn.getAttribute('id'))
            if (el[0] == btn.getAttribute('id')) {
                const xhr = new XMLHttpRequest()
                xhr.open("GET", "http://localhost:8080/data.html/:" + el[0])
                xhr.send()
                xhr.onload = ()=> {
                    if (xhr.readyState == 4 && xhr.status == 200) {
						document.getElementById('fields').innerHTML
						.replace('{{ .data }}', xhr.responseText)
					}
                }
            }
        }
    })
}

// createBtn.addEventListener('click', ()=> {
// 	document.getElementById('create_fields').style.display = 'flex';
//     // console.log('shit')
// })
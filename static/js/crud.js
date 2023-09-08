const crudButtons = document.getElementsByClassName('CRUD_buttons')
const crudMap = new Map()



crudMap.set('create_btn', 'create_fields')
crudMap.set('read_btn', 'read_fields')
crudMap.set('delete_btn', 'delete_fields')
crudMap.set('update_btn', 'update_fields')


for (const btn of crudButtons) {
    btn.addEventListener('click', () => {
        for (const el of crudMap) {
            if (el[0] == btn.getAttribute('id')) {
                document.getElementById(el[1]).style.display = 'flex';
                const xhr = new XMLHttpRequest()
                xhr.open("GET", "http://localhost:8080/data.html/:create")
                xhr.send()
                xhr.onload = ()=> {
                    if (xhr.readyState == 4 && xhr.status == 200) {
                        const data = xhr.response;
                        console.log(data);
                      } else {
                        console.log(`Error: ${xhr.status}`);
                      }
                }
            } else {
                document.getElementById(el[1]).style.display = 'none';
            }
        }
    })
}

// createBtn.addEventListener('click', ()=> {
// 	document.getElementById('create_fields').style.display = 'flex';
//     // console.log('shit')
// })
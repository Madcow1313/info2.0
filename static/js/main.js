const btn = document.getElementById('btn')
const aboutBtn = document.getElementById('about_btn')
const homePage = '/'
const about = '/about.html'

btn.addEventListener('click', ()=> {
	location.href = homePage
	console.log('wtf')
})

aboutBtn.addEventListener('click', ()=> {
	location.href = aboutBtn.getAttribute('href')
})
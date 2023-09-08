const hp = document.getElementById('homepage_btn');
const aboutBtn = document.getElementById('about_btn');
const homePage = '/';
const dataBtn = document.getElementById('data_btn');

const mainBtns = document.getElementsByClassName('buttons');

hp.addEventListener('click', ()=> {
	location.href = homePage
})

for (const el of mainBtns) {
	el.addEventListener('click', () => {
		location.href = el.getAttribute('href')
	});
};
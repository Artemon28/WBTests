// const order = document.querySelector('.order')
function renderOrder(uid) {
    fetch('/order/', {method: 'GET', body: uid})
        .then(response => response.json())
        .then(json => {
            order.innerHTML = `
                ${json.map((item, index) => {
                return `<div>${item.name}</div>`;
            }).join(' ')}`;
        });
}
document.querySelector('form').addEventListener('submit', (event) => {
    event.preventDefault();
    let uid = event.target.elements['id'].value;
    renderOrder(uid);
    // fetch('/order/', {method: 'GET'})
    //     .then(() => {
    //         renderOrder();
    //         event.target.reset();
    //     })
});
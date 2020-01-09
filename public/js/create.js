var socket = io('http://localhost:3000');

socket.on('connect', function(){
    socket.emit('requestDbNames');//Get database names to display to user
});

socket.on('gameNamesData', function(data){
    for(var i = 0; i < Object.keys(data).length; i++){
        
        var div = document.getElementById('game-list');
        var col = document.createElement('div');
        var button = document.createElement('button');
        col.classList.add("col-sm");

        var image = document.createElement('img');
        image.classList.add("image");

        button.setAttribute('id','gameButton');       
        button.innerHTML = data[i].name;

        image.src = data[i].image        
        image.setAttribute('onClick', "startGame('" + data[i].id + "')");
        col.appendChild(image);

        // button.className = "list-group-item list-group-item-primary";
        button.setAttribute('onClick', "startGame('" + data[i].id + "')");
        col.appendChild(document.createElement('br'));
        col.appendChild(button);

        div.appendChild(col);
        div.appendChild(document.createElement('br'));
    }
});

function startGame(data){
    window.location.href="/host/" + "?id=" + data;
}

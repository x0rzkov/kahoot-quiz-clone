var socket = io('https://api.eedama.org');
var playerAnswered = false;
var correct = false;
var name;
var score = 0;

var params = jQuery.deparam(window.location.search); //Gets the id from url

socket.on('connect', function() {
    //Tell server that it is host connection from game view
    socket.emit('player-join-game', params);

    document.getElementById('answer1').style.visibility = "visible";
    document.getElementById('answer2').style.visibility = "visible";
    document.getElementById('answer3').style.visibility = "visible";
    document.getElementById('answer4').style.visibility = "visible";
});

socket.on('noGameFound', function() {
    window.location.href = '../../'; //Redirect user to 'join game' page 
});

function answerSubmitted(num) {
    if (playerAnswered == false) {
        playerAnswered = true;

        socket.emit('playerAnswer', num); //Sends player answer to server

        //Hiding buttons from user
        document.getElementById('answer1').style.visibility = "hidden";
        document.getElementById('answer2').style.visibility = "hidden";
        document.getElementById('answer3').style.visibility = "hidden";
        document.getElementById('answer4').style.visibility = "hidden";
        document.getElementById('message').style.display = "block";
        document.getElementById('message').innerHTML = "Esperando que los demás participantes respondan...";

    }
}

//Get results on last question
socket.on('answerResult', function(data) {
    if (data == true) {
        correct = true;
    }
});

socket.on('questionOver', function(data) {
    if (correct == true) {
        document.body.style.backgroundColor = "#4CAF50";
        document.getElementById('message').style.display = "block";
        document.getElementById('message').innerHTML = "Correcto!";
    } else {
        document.body.style.backgroundColor = "#f94a1e";
        document.getElementById('message').style.display = "block";
        document.getElementById('message').innerHTML = "Incorrecto!";
    }
    document.getElementById('answer1').style.visibility = "hidden";
    document.getElementById('answer2').style.visibility = "hidden";
    document.getElementById('answer3').style.visibility = "hidden";
    document.getElementById('answer4').style.visibility = "hidden";
    socket.emit('getScore');
});

socket.on('newScore', function(data) {
    document.getElementById('scoreText').innerHTML = "Score: " + data;
});

socket.on('nextQuestionPlayer', function(data) {
    correct = false;
    playerAnswered = false;

    console.log('playerGame.gameQuestions', data)

    document.getElementById('question').innerHTML = data.q1;

    if (typeof data.i1 !== 'undefined') {
        document.getElementById('banner').src = data.i1;
        document.getElementById('banner').style.display = "block";
    } else {
        document.getElementById('banner').style.display = "none";
    }

    if (typeof data.a1 !== 'undefined') {
        document.getElementById('answer1').innerHTML = data.a1;
        document.getElementById('answer1').style.display = "block";
    } else {
        document.getElementById('answer1').style.display = "none";
    }

    if (typeof data.a2 !== 'undefined') {
        document.getElementById('answer2').innerHTML = data.a2;
        document.getElementById('answer2').style.display = "block";
    } else {
        document.getElementById('answer2').style.display = "none";
    }

    if (typeof data.a3 !== 'undefined') {
        document.getElementById('answer3').innerHTML = data.a3;
        document.getElementById('answer3').style.display = "block";
    } else {
        document.getElementById('answer3').style.display = "none";
    }

    if (typeof data.a4 !== 'undefined') {
        document.getElementById('answer4').innerHTML = data.a4;
        document.getElementById('answer4').style.display = "block";
    } else {
        document.getElementById('answer4').style.display = "none";
    }

    var correctAnswer = data.correct;

    document.getElementById('answer1').style.visibility = "visible";
    document.getElementById('answer2').style.visibility = "visible";
    document.getElementById('answer3').style.visibility = "visible";
    document.getElementById('answer4').style.visibility = "visible";
    document.getElementById('message').style.display = "none";
    document.body.style.backgroundColor = "white";

});

socket.on('hostDisconnect', function() {
    window.location.href = "../../";
});

socket.on('playerGameData', function(data) {
    for (var i = 0; i < data.length; i++) {
        if (data[i].playerId == socket.id) {
            document.getElementById('nameText').innerHTML = "Nombre: " + data[i].name;
            document.getElementById('scoreText').innerHTML = "Puntaje: " + data[i].gameData.score;
        }
    }
});

socket.on('gameQuestions', function(data) {

    console.log('playerGame.gameQuestions', data)

    document.getElementById('question').innerHTML = data.q1;

    if (typeof data.i1 !== 'undefined') {
        document.getElementById('banner').src = data.i1;
        document.getElementById('banner').style.display = "block";
    } else {
        document.getElementById('banner').style.display = "none";
    }

    if (typeof data.a1 !== 'undefined') {
        document.getElementById('answer1').innerHTML = data.a1;
        document.getElementById('answer1').style.display = "block";
    } else {
        document.getElementById('answer1').style.display = "none";
    }

    if (typeof data.a2 !== 'undefined') {
        document.getElementById('answer2').innerHTML = data.a2;
        document.getElementById('answer2').style.display = "block";
    } else {
        document.getElementById('answer2').style.display = "none";
    }

    if (typeof data.a3 !== 'undefined') {
        document.getElementById('answer3').innerHTML = data.a3;
        document.getElementById('answer3').style.display = "block";
    } else {
        document.getElementById('answer3').style.display = "none";
    }

    if (typeof data.a4 !== 'undefined') {
        document.getElementById('answer4').innerHTML = data.a4;
        document.getElementById('answer4').style.display = "block";
    } else {
        document.getElementById('answer4').style.display = "none";
    }

    var correctAnswer = data.correct;

});

socket.on('GameOver', function() {
    document.body.style.backgroundColor = "#FFFFFF";
    document.getElementById('question').style.visibility = "hidden";
    document.getElementById('answer1').style.visibility = "hidden";
    document.getElementById('answer2').style.visibility = "hidden";
    document.getElementById('answer3').style.visibility = "hidden";
    document.getElementById('answer4').style.visibility = "hidden";
    document.getElementById('message').style.display = "block";
    document.getElementById('message').innerHTML = "GAME OVER";
});
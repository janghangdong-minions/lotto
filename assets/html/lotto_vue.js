var app = new Vue({
    el: '#app',
    data: {
        message: '로토번호 생성기',
        numbers: [21, 54, 64, 21, 5, 5, 235] //요놈을 재배치
    },
    methods: {
        changeNumber: function() {
            this.numbers = getRandomNums()
                //this.numbers = getRandomNums() //7개의 로토넘버
                //await this.setNumberColor()
                //await setNumberColor() //동기식 ->비동기
        }
    },
    watch: {
        numbers: async() => {
            await this.setNumberColor()
        }
    }

})



function getRandomNums() {
    var lotto = []
    for (var i = 0; i < 7; i++) {
        var num = Math.floor(Math.random() * 44) + 1;
        lotto.push(num);
    }
    return lotto
}


//범위에 따라 특정컬러로 엘리먼트의 색을 바꾼다.
function setNumberColor() {
    numObj = document.getElementsByClassName('circle');

    for (let obj = 0; obj < numObj.length; obj++) {
        num = numObj[obj].textContent
        console.log(num);
        if (num <= 10) {
            numObj[obj].style.color = 'yellow';
        } else if (num > 10 && num <= 20) {
            numObj[obj].style.color = 'blue';
        } else if (num > 20 && num <= 30) {
            numObj[obj].style.color = 'red';
        } else if (num > 30 && num <= 40) {
            numObj[obj].style.color = 'black';
        } else if (num > 40 && num <= 50) {
            numObj[obj].style.color = 'green';
        }
    }
}
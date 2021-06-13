//lotto js 파일
function setNumberColor(){
    numObj = document.getElementsByClassName('circle');
    
    for(let i=0; i<numObj.length; i++){
        num = numObj[i].textContent
        if (num<=10){
            numObj[i].style.color = 'yellow';
        }else if(num>10 && num<=20){
            numObj[i].style.color = 'blue';
        }else if(num>20 && num<=30){
            numObj[i].style.color = 'red';
        }else if(num>30 && num<=40){
            numObj[i].style.color = 'black';
        }else if(num>40 && num<=50){
            numObj[i].style.color = 'green';
        }
    }
}
//로토번호들의 컬러를바꾼다.
setNumberColor();
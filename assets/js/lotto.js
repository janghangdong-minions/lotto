//lotto js 파일
function setNumberColor(){
    numObj = document.getElementsByClassName('circle');
    
    for(let obj=0; obj<numObj.length; obj++){
        num = numObj[obj].textContent
        if (num<=10){
            numObj[obj].style.color = 'yellow';
        }else if(num>10 && num<=20){
            numObj[obj].style.color = 'blue';
        }else if(num>20 && num<=30){
            numObj[obj].style.color = 'red';
        }else if(num>30 && num<=40){
            numObj[obj].style.color = 'black';
        }else if(num>40 && num<=50){
            numObj[obj].style.color = 'green';
        }
    }
}
//로토번호들의 컬러를바꾼다.
setNumberColor();
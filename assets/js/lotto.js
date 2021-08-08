//lotto js 파일

//범위에 따라 특정컬러로 엘리먼트의 색을 바꾼다.
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

//로토넘버를 재생성해주는 함수
function changeNumber (){
    fetch('/rand') //[url,option] 옵션값을 입력하지 않은면 GET 매소드로 간주한다.
    .then((response) => response.json())
    .then((data) => {
        return data
    })
    .then(nums => {
        mainNumber = document.getElementById("main");
        mainNumber.innerHTML="";
        console.log(nums)
        for (let index = 0; index < nums.BasicNums.length; index++) {
            mainNumber.innerHTML += `<span class="circle">${nums.BasicNums[index]}</span>`;
        }
        setNumberColor();
    })

}
//로토번호들의 컬러를바꾼다.
setNumberColor();
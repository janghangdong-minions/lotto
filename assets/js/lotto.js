//lotto js 파일
const LS = "numSet";
let lsNums = [];
const ol = document.getElementById("favNum");

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
    .then(function(nums){
        mainNumber = document.getElementById("main");
        mainNumber.innerHTML="";
        console.log(nums)
        for (let index = 0; index < nums.BasicNums.length; index++) {
            mainNumber.innerHTML += `<span class="circle">${nums.BasicNums[index]}</span>`;
        }
        setNumberColor();
    })

}

//원하는 로토넘버를 로컬스토리지에 저장한다.
function putNumber (){
    ol.innerHTML='';  //리스트초기화
    let numSets = [];
    let numDiv = document.getElementById("main");
    numSpans = numDiv.getElementsByClassName("circle");
    for (index = 0; index < numSpans.length; index++){
        numSets.push(numSpans[index].textContent);
    }
    const newId = lsNums.length + 1;
    const numSetObj = {
        id : newId,
        num : numSets
    }
    lsNums.push(numSetObj);
    localStorage.setItem(LS,JSON.stringify(lsNums));
    loadLS()
}



//로컬스토리지에 담은 로토넘버를 보여준다.
function loadLS (){  
    let s = localStorage.getItem(LS);
    if(s !== null ){
        const parsedNumSet = JSON.parse(s);
        //lsNums.push(parsedNumSet);
        parsedNumSet.forEach(function(numSet){
            showFavNum(numSet.num);
        });
    }
}

function loadLS2 (){
    let v = localStorage.getItem(LS);
    if(v !== null ){
        const parsedNumSet = JSON.parse(s);
        lsNums.push(parsedNumSet);
    }
}



//로컬스토리지에서 불러온 넘버를 렌더한다.
function showFavNum(num){
    let li = document.createElement("li");
    for (index = 0; index < num.length; index++){
        li.innerHTML += `<span class="circle">${num[index]}</span>`;
    }
    ol.appendChild(li);
    setNumberColor();
}


//로토번호들의 컬러를바꾼다.
setNumberColor();
//로컬스토리지 로드
loadLS2();
loadLS();
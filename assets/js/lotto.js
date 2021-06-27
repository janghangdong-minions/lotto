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

var nums = 0;

function rendomData() {
    $.ajax({
        url: "/rend" ,
        type: "GET",
        accepts: {
          mycustomtype: 'application/json'
        },
        success:function(data){
            console.log(data);
            let addDiv = document.getElementById("addRend");
            console.log(addDiv);
            let addSpan = document.getElementById(`rend_${nums}`);
            if (!addSpan){
                addDiv.innerHTML += `<div id="rend_${nums}">`
                for (let index = 0; index < nums; index++) {
                    addDiv.innerHTML += `<span class="circle">${data.BasicNums[0]}</span>`;
                    addDiv.innerHTML += `<span class="circle">${data.BasicNums[1]}</span>`;
                    addDiv.innerHTML += `<span class="circle">${data.BasicNums[2]}</span>`;
                    addDiv.innerHTML += `<span class="circle">${data.BasicNums[3]}</span>`;
                    addDiv.innerHTML += `<span class="circle">${data.BasicNums[4]}</span>`;
                    addDiv.innerHTML += `<span class="circle">${data.BasicNums[5]}</span>`;
                    addDiv.innerHTML += `<span class="circle">${data.BasicNums[6]}</span>`;
                }
                addDiv.innerHTML += `</div><br />`
                nums +=1;
            }
        }
      });
};
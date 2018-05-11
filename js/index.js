function submit(){
    var userName=$("#username").val();
    var password=$("#pass").val();
    var mail=$("#mail").val();
    if (mail.length<7 || !mail.includes("@")){
        alert("پست الکترونیکی خود را به درستی وارد کنید");
        
    }else{

        var xhttp = new XMLHttpRequest();
        var datas ="user="+userName+"&pass="+password+"&mail="+mail
        xhttp.onreadystatechange = function() {
          if (this.readyState == 4 && this.status == 200) {
      
            var resp=this.responseText.split("&");
            if(resp[0]=="1"){
                login();
            }else{
                alert("مجددا تلاش کنید");
            }
             
           
          }
        };
      
        xhttp.open("POST", "http://localhost:3000/submit", true);
        xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xhttp.send(datas);
      
    }             
 

}


function login(){
    var userName=$("#username").val();
    var password=$("#pass").val();
    
           
  var xhttp = new XMLHttpRequest();
  var datas ="user="+userName+"&pass="+password;
  xhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {

        var resp=this.responseText.split("&");
        if(resp[0]=="1"){

            localStorage.setItem("pressUser",resp[1]);
            $("#profile").text(resp[1]+"خوش آمدید");
            $("#login").hide(100);
        }else{
            localStorage.removeItem("pressUser");
            alert("خطا...مجددا تلاش کنید");
        }
       

       
     
    }
  };

  xhttp.open("POST", "http://localhost:3000/login", true);
  xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhttp.send(datas);


}






function showSubmit(){
    $("#mail").show(200);
    $("#showSubmit").hide();
    $("#logon").css("background-color","blue");
    $("#logon").attr("onclick","submit()");
}
  
function increaseValue(id) {
  var ID="#"+id
  var value = parseInt($(ID).text(), 10);
  value = isNaN(value) ? 0 : value;
  value++;
  $(ID).text("+"+value) ;
}

function decreaseValue(id) {
  var ID="#"+id  
  var value =parseInt($(ID).text(), 10);
  value = isNaN(value) ? 0 : value;
  value < 1 ? value = 1 : '';
  value--;
  $(ID).text("+"+value) ;
}


function showleaf(){
  
  $('#dialogue').click(function(){
    $('#leaf').show(500);
});
 }

 function clsleaf() {
   
  $('#clsbox').click(function(){
    $('#leaf').hide(500);
 });
 }


function getLastNews(){

    var user=localStorage.getItem("pressUser");
    var datas;
    if( user !=''){

        datas="user="+user+"&key=1";
    }else{
        datas="user=0&key=0";
    }



    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {

        var resp=this.responseText;

        var division = document.createElement("div");
        $(division).attr("id","added");

        division.innerHTML=resp;
       document.getElementById("login").after(division);
  

     
    }
  };

  xhttp.open("POST", "http://localhost:3000/getnews", true);
  xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhttp.send(datas);



}


$(document).ready(function(){
  
 
    
    $('#leaf').hide();
    $('#splash').fadeOut(2000);

    getLastNews();







var elementPosition = $('.navbar').offset();
var chatposition = $('#dialogue').offset();

$(window).scroll(function(){
        if($(window).scrollTop() > elementPosition.top){
              $('.navbar').css('opacity','0.8');
              
        } else {
            $('.navbar').css('opacity','1');
            
        }    


      });


///////////////////////////////////
var sundte = new Date
var yeardte = sundte.getFullYear();
var monthdte = sundte.getMonth();
var dtedte = sundte.getDate();
var daydte = sundte.getDay();
var sunyear
switch (daydte) {
    case 0:
        var today = "يکشنبه";
        break;
    case 1:
        var today = "دوشنبه";
        break;
    case 2:
        var today = "سه شنبه";
        break;
    case 3:
        var today = "چهارشنبه";
        break;
    case 4:
        var today = "پنچشنبه";
        break;
    case 5:
        var today = "جمعه";
        break;
    case 6:
        var today = "شنبه";
        break;
}
switch (monthdte) {
    case 0:

        sunyear = yeardte - 622;
        if (dtedte <= 20) {
            var sunmonth = "دي";
            var daysun = dtedte + 10;
        } else {
            var sunmonth = "بهمن";
            var daysun = dtedte - 20;
        }
        break;
    case 1:

        sunyear = yeardte - 622;
        if (dtedte <= 19) {
            var sunmonth = "بهمن";
            var daysun = dtedte + 11;
        } else {
            var sunmonth = "اسفند";
            var daysun = dtedte - 19;
        }
        break;
    case 2:
        {
            if ((yeardte - 621) % 4 == 0)
                var i = 10;
            else
                var i = 9;
            if (dtedte <= 20) {
                sunyear = yeardte - 622;
                var sunmonth = "اسفند";
                var daysun = dtedte + i;
            } else {
                sunyear = yeardte - 621;
                var sunmonth = "فروردين";
                var daysun = dtedte - 20;
            }
        }
        break;
    case 3:

        sunyear = yeardte - 621;
        if (dtedte <= 20) {
            var sunmonth = "فروردين";
            var daysun = dtedte + 11;
        } else {
            var sunmonth = "ارديبهشت";
            var daysun = dtedte - 20;
        }
        break;
    case 4:

        sunyear = yeardte - 621;
        if (dtedte <= 21) {
            var sunmonth = "ارديبهشت";
            var daysun = dtedte + 10;
        } else {
            var sunmonth = "خرداد";
            var daysun = dtedte - 21;
        }

        break;
    case 5:

        sunyear = yeardte - 621;
        if (dtedte <= 21) {
            var sunmonth = "خرداد";
            var daysun = dtedte + 10;
        } else {
            var sunmonth = "تير";
            var daysun = dtedte - 21;
        }
        break;
    case 6:

        sunyear = yeardte - 621;
        if (dtedte <= 22) {
            var sunmonth = "تير";
            var daysun = dtedte + 9;
        } else {
            var sunmonth = "مرداد";
            var daysun = dtedte - 22;
        }
        break;
    case 7:

        sunyear = yeardte - 621;
        if (dtedte <= 22) {
            var sunmonth = "مرداد";
            var daysun = dtedte + 9;
        } else {
            var sunmonth = "شهريور";
            var daysun = dtedte - 22;
        }
        break;
    case 8:

        sunyear = yeardte - 621;
        if (dtedte <= 22) {
            var sunmonth = "شهريور";
            var daysun = dtedte + 9;
        } else {
            var sunmonth = "مهر";
            var daysun = dtedte + 22;
        }
        break;
    case 9:

        sunyear = yeardte - 621;
        if (dtedte <= 22) {
            var sunmonth = "مهر";
            var daysun = dtedte + 8;
        } else {
            var sunmonth = "آبان";
            var daysun = dtedte - 22;
        }
        break;
    case 10:

        sunyear = yeardte - 621;
        if (dtedte <= 21) {
            var sunmonth = "آبان";
            var daysun = dtedte + 9;
        } else {
            var sunmonth = "آذر";
            var daysun = dtedte - 21;
        }

        break;
    case 11:

        sunyear = yeardte - 621;
        if (dtedte <= 19) {
            var sunmonth = "آذر";
            var daysun = dtedte + 9;
        } else {
            var sunmonth = "دي";
            var daysun = dtedte + 21;
        }
        break;

}
document.getElementById("time").innerHTML = "امروز "+today+"&nbsp;"+daysun+"&nbsp;"+sunmonth+"&nbsp;"+sunyear;


    



    });


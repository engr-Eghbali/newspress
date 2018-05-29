function showLogin(){
    $("#search").hide();
    $("#cat").hide();
    $("#username").show();
    $("#password").show();
    $("#submitLink").show();
    $("#form").show();
}

function showSubmit(){
    $("#mail").show();
    $("#submitLink").hide();
    $("#submition").attr("onclick","submit()");
}

function closeForm(){
    $("#newsForm").hide();
}
function newsForm(){
    $("#newsForm").show();
}
function clsCm(){
    $("#showNews").hide();
}



function submit(){
    var userName=$("#username").val();
    var password=$("#password").val();
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
                login(userName,password);
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


function login(u,p){
    var userName=u;
    var password=p;
    
    if(userName==null || password==null){
        return
    }
           
  var xhttp = new XMLHttpRequest();
  var datas ="user="+userName+"&pass="+password;
  xhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {

        var resp=this.responseText.split("&");
        if(resp[0]=="1"){

            localStorage.setItem("pressUser",resp[1]);
            localStorage.setItem("pressPass",password);
            $("#profile").text(resp[1]+"خوش آمدید  ");
            $("#loginSubmit").text("خروج");
            $("#loginSubmit").attr("onclick","logout()");
            $("#loginSubmit").css({"color":"tomato"});

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

function doLogin(){
    var u=$("#username").val();
    var p=$("#password").val();
    login(u,p);
}


function logout(){
    localStorage.removeItem("pressUser");
    localStorage.removeItem("pressPass");
    location.reload();
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
        $(division).attr("id","ourserv");

        division.innerHTML=resp;
       document.getElementById("line").after(division);
  

     
    }
  };

  xhttp.open("POST", "http://localhost:3000/getnews", true);
  xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhttp.send(datas);



}




//////////////////////////////////////////////////////
$(document).ready(function(){
  
    
    if(localStorage.getItem("pressUser")!=""){
        login(localStorage.getItem("pressUser"),localStorage.getItem("pressPass"));
    }


    getLastNews();



});



function showSearch(){

    $("#submitLink").hide();
    $("#username").hide();
    $("#password").hide();
    $("#mail").hide();
    $("#search").show();
    $("#form").show();
    $("#cat").show();
    $("#submition").attr("onclick","search()");

}
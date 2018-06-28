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
            //$("#profile").text(resp[1]+"خوش آمدید  ");
            $("#loginSubmit").text("خروج");
            $("#loginSubmit").attr("onclick","logout()");
            $("#loginSubmit").css({"color":"tomato"});
            $("#form").hide();

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


function more(id){

    var user=0;
    if (localStorage.getItem("pressUser")!=""){
        user=localStorage.getItem("pressUser");
    }
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {

        var resp=this.responseText;
        document.getElementById("post").innerHTML=resp;
        $("#showNews").show();
     
    }
  };

  xhttp.open("POST", "http://localhost:3000/more", true);
  xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhttp.send("newsID="+id+"&pressUser="+user);


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


function like(id){
      

  var username=localStorage.getItem("pressUser");  
  if  (username ==null){
      alert("ابتدا باید ثبت نام کنید یا وارد شوید");

  }else{

    var newsID=id;
  
    var xhttp = new XMLHttpRequest();
    var datas ="username="+username+"&newsID="+newsID;
    xhttp.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
  
          var resp=this.responseText;
          if(resp=="1"){
  
            
                     var cssID="count";
                     var value = parseInt(document.getElementById(cssID).innerHTML);
                     
                     value = isNaN(value) ? 0 : value;
                     value++;
                     document.getElementById(cssID).innerHTML="";
                     document.getElementById("likes").innerHTML="<i class=\"fas fa-heart\"></i>";
                     document.getElementById(cssID).innerHTML=value;
                     $("#likes").attr("onclick","dislike('"+newsID+"')");
          }
         
  
         
       
      }

  }

  xhttp.open("POST", "http://localhost:3000/like", true);
  xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhttp.send(datas);

 
}
}



function dislike(id){
      
    var username=localStorage.getItem("pressUser");  
    if  (username ==""){
        alert("ابتدا باید ثبت نام کنید یا وارد شوید");
  
    }else{
  
      var newsID=id;
    
      var xhttp = new XMLHttpRequest();
      var datas ="username="+username+"&newsID="+newsID;
      xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
    
            var resp=this.responseText;
            if(resp=="1"){
    
              
                var cssID="count";
                var value = parseInt(document.getElementById(cssID).innerHTML);
                
                value = isNaN(value) ? 0 : value;

                       value < 1 ? value = 1 : '';
                       value--;
                       document.getElementById(cssID).innerHTML="";
                       document.getElementById("likes").innerHTML="<i class=\"far fa-heart\"></i>";
                       document.getElementById(cssID).innerHTML=value;
                       $("#likes").attr("onclick","like('"+newsID+"')");
  
                     
            }
           
    
           
         
        }
  
    }
  
    xhttp.open("POST", "http://localhost:3000/dislike", true);
    xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhttp.send(datas);
  
   
  }
  }


  function sendCm(newsID){
      var user=localStorage.getItem("pressUser");
      if(user==null){
          alert("ابتدا باید وارد شوید یا ثبت نام کنید.");
          return;
      }
      var Cm=$("#inputCm").val();
      if(Cm.length !=""){


        var xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
    
            if(this.responseText!="0"){

                var resp=this.responseText;
                var div=document.createElement("div");
                div.className="cm";
                div.innerHTML=resp;
                $("#cms").append(div);
    
            }else{
                alert("خطا، مجددا تلاش کنید");
            }
        }
      };
    
      xhttp.open("POST", "http://localhost:3000/Cm", true);
      xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
      xhttp.send("newsID="+newsID+"&username="+user+"&Cm="+Cm);


      }else{
          alert("پیغام نمی تواند خالی باشد ");
      }
  }




  function uploader(id) {
    var mFile = document.getElementById('newsFile');
    var photo = mFile.files[0];
    
    var fd = new FormData();

    fd.append("name",id);
    fd.append("path","./news");
    fd.append("image",photo);


    if (photo.size > 520000){
        alert("picture size must be lower than 500kb");
    }else {
        $.ajax({
            url: 'http://localhost:3000/upload',
            type: "POST",
            enctype: 'multipart/form-data',
            processData: false,
            contentType: false,
            cache: false,
            timeout: 600000,
            data: fd,

            success: function (data) {
                alert("خبر شما درج شد");
                location.reload();
            },
            error: function () {
                alert("خبر شما درج شد");
                location.reload();
            }
        });
    }

}







  function sendNews(){

      var title=$("#newsTitle").val();
      var keywords=$("#newsCategory").val();
      var text=$("#newsText").val();

      if(text.length<300){

        alert("متن خبر باید حداقل 300 کاراکتر باشد");
        return;
      }

      var user=localStorage.getItem("pressUser");
      var dt = new Date();
      var stamp=(("0"+dt.getDate()).slice(-2)) +"."+ (("0"+(dt.getMonth()+1)).slice(-2)) +"."+ (dt.getFullYear()) +" "+ (("0"+dt.getHours()).slice(-2)) +":"+ (("0"+dt.getMinutes()).slice(-2));

      if(title=="" || keywords=="" || text==""){
          alert("پر کردن تمامی فیلدها اجباری میباشد");
      }else{


        
        var xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
    
            if(this.responseText!="0"){

                uploader(this.responseText);
    
            }else{
                alert("خطا، مجددا تلاش کنید");
            }
        }
      };
    
      xhttp.open("POST", "http://localhost:3000/sendNews", true);
      xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
      xhttp.send("username="+user+"&title="+title+"&keywords="+keywords+"&time="+stamp+"&text="+text);


      }

  }


function search(){
    var category=$("#cat").val();
    var query   =$("#search").val();

    if (query.length<2){
        alert("مقدار جست و جو نمی تواند خالی باشد.");
    }else{

        var xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
    
            if(this.responseText!="0"){

                $("#slider").remove();
                $(".nivo-controlNav").remove();
                $("#title").text("نتایج یافت شده");
                $("#ourserv").remove();
               
               var resp=this.responseText;
               
               var division = document.createElement("div");
               $(division).attr("id","ourserv");
       
               if(resp==""){
                   resp="<div id=\"nf404\"><p1>متاسفانه نتیجه ای یافت نشد.</p1><br><br><img src=\"./images/sad.svg\" id=\"sad\"><br><p2>برای نتایج بهتر از کلمات کلیدی برای جست و جو استفاده کنید</p2></div>";
               }
               division.innerHTML=resp;
              document.getElementById("line").after(division);
       
    
            }else{
                alert("خطا، مجددا تلاش کنید");
            }
        }
      };
    
      xhttp.open("POST", "http://localhost:3000/searchNews", true);
      xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
      xhttp.send("cat="+category+"&query="+query);


      }

    }


    function readURL(input) {
        if (input.files && input.files[0]) {
            var reader = new FileReader();

            reader.onload = function (e) {
                $('#preview')
                    .attr('src', e.target.result)
                    .width(200)
                    .height(200);
            };
             
            $("#preview").show();
            reader.readAsDataURL(input.files[0]);
        }
    }



function editNews(id){


    var title=$("#newsTitle").val();
    var keywords=$("#newsCategory").val();
    var text=$("#newsText").val();

    if(text.length<300){

      alert("متن خبر باید حداقل 300 کاراکتر باشد");
      return;
    }

    var user=localStorage.getItem("pressUser");
    var dt = new Date();
    var stamp=(("0"+dt.getDate()).slice(-2)) +"."+ (("0"+(dt.getMonth()+1)).slice(-2)) +"."+ (dt.getFullYear()) +" "+ (("0"+dt.getHours()).slice(-2)) +":"+ (("0"+dt.getMinutes()).slice(-2));

    if(title==""  || text==""){
        alert("پر کردن تمامی فیلدها اجباری میباشد");
    }else{


      
      var xhttp = new XMLHttpRequest();
      xhttp.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
  
          if(this.responseText!="0"){

            var mFile = document.getElementById('newsFile');
           // var photo = mFile.files[0];
            if(mFile.value !=""){

                uploader(this.responseText);
            }else{
                alert("خبر با موفقیت ویرایش شد");
                location.reload();
            }       
              
  
          }else{
              alert("خطا، مجددا تلاش کنید");
          }
      }
    };
  
    xhttp.open("POST", "http://localhost:3000/editNews", true);
    xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhttp.send("username="+user+"&title="+title+"&id="+id+"&keywords="+keywords+"&time="+stamp+"&text="+text);


    }

}    



function edit(id){

    if(localStorage.getItem("pressUser")==null){
        alert("ابتدا باید ثبت نام کنید یا وارد شوید");
        return;
    }
    var text=$(".textCm").text();
    text=text.replace(text.substring(text.indexOf("@"),text.indexOf(":")),"");
    var title=$("#head"+id).text();
    var img=$("#img"+id).attr('src');
    $("#newsTitle").val(title);
    $("#newsText").val(text);
    $("#preview").attr('src',img);
    $("#sendNews").attr("onclick","editNews('"+id+"')");
    newsForm();
    $("#showNews").hide();
}



$(document).ready(function(){


 
    var input = document.getElementById('uploadImage');
   input.onchange = function(evt){
       var tgt = evt.target || window.event.srcElement, 
           files = tgt.files;
  
       if (FileReader && files && files.length) {
           var fr = new FileReader();
           fr.onload = function () {
               localStorage['pressProfilePic'] = fr.result;
           }
           fr.readAsDataURL(files[0]);
       
       location.reload();
         }
       
   }
   if (localStorage['pressProfilePic'].length >100){
    var el = document.getElementById('profilepic');
    el.style.backgroundImage = 'url(' + localStorage['raadifeProfilePic'] + ')';
   }
   
  });
  
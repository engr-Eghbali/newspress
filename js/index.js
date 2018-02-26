function increaseValue() {
  var value = parseInt(document.getElementById('number').value, 10);
  value = isNaN(value) ? 0 : value;
  value++;
  document.getElementById('number').value = value;
}

function decreaseValue() {
  var value = parseInt(document.getElementById('number').value, 10);
  value = isNaN(value) ? 0 : value;
  value < 1 ? value = 1 : '';
  value--;
  document.getElementById('number').value = value;
}

$(document).ready(function(){


var elementPosition = $('.navbar').offset();

$(window).scroll(function(){
        if($(window).scrollTop() > elementPosition.top){
              $('.navbar').css('opacity','0.8');
              
        } else {
            $('.navbar').css('opacity','1');
            
        }    
});

});


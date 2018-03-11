(function(root, document) {
  'use strict';
  
  var message = document.getElementById('message');
  var button = document.getElementById('button');
  
  if (message && button) {
    var clickHandler = function(event) {
      var xhr = new XMLHttpRequest({mozSystem: true});
      xhr.open("GET", "/ajax/param", true);
      xhr.onreadystatechange = function () {
        if (xhr.status === 200 && xhr.readyState === 4) {
          message.innerHTML = "Ajax succeded";
        }
      };

      xhr.onerror = function () {
        message.innerHTML = "Ajax failed.";
      };
      
      xhr.send();
    };

    button.addEventListener('click', clickHandler);
  }
  
  var message2 = document.getElementById('message2');
  var button2 = document.getElementById('button2');
  
  if (message && button) {
    var clickHandler = function(event) {
      var xhr = new XMLHttpRequest({mozSystem: true});
      xhr.open("GET", "/ajax2/param", true);
      xhr.onreadystatechange = function () {
        if (xhr.status === 200 && xhr.readyState === 4) {
          message2.innerHTML = "Ajax succeded";
        }
      };

      xhr.onerror = function () {
        message2.innerHTML = "Ajax failed.";
      };
      
      xhr.send();
    };

    button2.addEventListener('click', clickHandler);
  }

}(this, this.document));

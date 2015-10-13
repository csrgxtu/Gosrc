<html>
  <head>
    <title>Auth0 Testing</title>

    <script src="/static/js/jquery.min.js"></script>
    <!-- Auth0 lock script -->
    <script src="//cdn.auth0.com/js/lock-7.9.min.js"></script>

    <!-- Setting the right viewport -->
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no" />
  </head>
  <body>
    <script>
      var lock = null;
      $(document).ready(function() {
       lock = new Auth0Lock('AFRZV4NudbDz0SDFT4qHUhKzDzbHZbti', 'csrgxtu.auth0.com');
      });

      var userProfile;

      $('.btn-login').click(function(e) {
        e.preventDefault();
        lock.show(function(err, profile, token) {
          if (err) {
            // Error callback
            alert('There was an error');
          } else {
            // Success callback

            // Save the JWT token.
            localStorage.setItem('userToken', token);

            // Save the profile
            userProfile = profile;
          }
        });
      });
    </script>
    <center>
      <h3> Here Testing Auth0 </h3>
      <input type="submit" class="btn-login" />
    </center>
  </body>
</html>

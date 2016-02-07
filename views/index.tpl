<div id="page-wrapper">
    <div class="container">
        <div class="row">
            <div class="col-md-8 col-md-offset-2">
                <h1>Welcome to Pratl</h1>
                <hr>
                {{if .Error}}
                    <div class="alert alert-danger">
                        {{.Error}}
                    </div>
                {{end}}
                <div class="row">
                    <div class="col-md-5">
                        <h4>Log In</h4>
                        <p class="text-muted">
                            Already a registered user?
                        </p>
                        <form method="post">
                            <input type="hidden" name="Action" value="login">
                            <div class="form-group">
                                <label for="email">Email</label>
                                <input type="email" class="form-control" id="email" name="Email" value="{{.LoginForm.Email}}">
                            </div>
                            <div class="form-group">
                                <label for="password">Password</label>
                                <input type="password" class="form-control" id="password" name="Password" value="{{.LoginForm.Password}}">
                            </div>
                            <button type="submit" class="btn btn-primary">
                                Log In
                            </button>
                        </form>
                        <br><br>
                    </div>
                    <div class="col-md-5 col-md-offset-1">
                        <h4>Register</h4>
                        <p class="text-muted">
                            Don't have an account?
                        </p>
                        <form method="post">
                            <input type="hidden" name="Action" value="register">
                            <div class="form-group">
                                <label for="email">Email</label>
                                <input type="email" class="form-control" id="email" name="Email" value="{{.RegisterForm.Email}}">
                            </div>
                            <div class="form-group">
                                <label for="name">Name</label>
                                <input class="form-control" id="name" name="Name" value="{{.RegisterForm.Name}}">
                            </div>
                            <div class="form-group">
                                <label for="password">Password</label>
                                <input type="password" class="form-control" id="password" name="Password" value="{{.RegisterForm.Password}}">
                            </div>
                            <div class="form-group">
                                <label for="password2">Confirm Password</label>
                                <input type="password" class="form-control" id="password2" name="Password2" value="{{.RegisterForm.Password2}}">
                            </div>
                            <button type="submit" class="btn btn-primary">
                                Register
                            </button>
                        </form>
                    </div>
                </div>
                <hr>
                <p class="text-muted">
                    Copyright 2016 - Nathan Osman
                </p>
            </div>
        </div>
    </div>
</div>

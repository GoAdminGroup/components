package theme1

var List = map[string]string{"login": `{{define "login"}}
    <!DOCTYPE html>
    <html>
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">

        <title>{{.Title}}</title>
        <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">

        <link rel="stylesheet" href="{{link .CdnUrl .UrlPrefix "/assets/login/dist/all.min.css"}}">

        <style>
            body.login-page {
                background-color: #2d3a4b;
                height: 80%;
            }

            div.login-logo a {
                color: white;
            }

            .text-center.text-muted {
                color: white;
            }

            .text-center.text-muted a {
                color: #92adce;
            }

            button.btn.btn-flat {
                background-color: #6a83a2;
            }

            .captcha {
                cursor: pointer;
                border: 1px #e6e6e6 solid;
            }
        </style>

    </head>
    <body class="hold-transition login-page" data-gr-c-s-loaded="true">
    <div class="login-box">
        <div class="login-logo">
            <a href="/"><b>{{.Title}}</b></a>
        </div>

        <div class="login-box-body">
            <form action="##" method="post" onsubmit="return false" id="sign-in-form">
                <div class="form-group has-feedback 1">
                    <input type="text" required class="form-control" placeholder="{{lang "username"}}" id="username">
                    <span class="glyphicon glyphicon-envelope form-control-feedback"></span>
                </div>
                <div class="form-group has-feedback 1">
                    <input type="password" required class="form-control" placeholder="{{lang "password"}}" id="password">
                    <span class="glyphicon glyphicon-lock form-control-feedback"></span>
                </div>
                {% if .CaptchaDigits %}
                    <div class="form-group has-feedback 1">
                        <div class="row">
                            <div class="col-xs-7">
                                <input type="text" class="form-control" placeholder="{{lang "captcha"}}" id="captcha">
                            </div>
                            <div class="col-xs-5">
                                <img class="captcha" src="{% .CaptchaImgSrc %}" alt="" width="110" height="34">
                            </div>
                        </div>
                        <input type="hidden" value="{% .CaptchaID %}" id="captcha_id">
                    </div>
                {% end %}
                <div class="row">
                    <div class="col-xs-8">
                    </div>
                    <div class="col-xs-4">
                        <button type="submit" class="btn btn-primary btn-block btn-flat" onclick="submitData()">{{lang "login"}}</button>
                    </div>
                </div>
            </form>

        </div>
    </div>

    <div class="text-center text-muted">
        <small>
            <strong>Powered by <a href="https://github.com/GoAdminGroup/go-admin"
                                  target="_blank">GoAdmin</a></strong>
        </small>
    </div>

    </body>

    {% if .TencentWaterProofWallData.AppID  %}
        <script src="https://ssl.captcha.qq.com/TCaptcha.js"></script>
    {% end %}

    <script src="{{link .CdnUrl .UrlPrefix "/assets/login/dist/all.min.js"}}"></script>
    <script>

        {% if .TencentWaterProofWallData.AppID  %}

        let captcha = new TencentCaptcha("{% .TencentWaterProofWallData.AppID %}", function (res) {
            console.log(res);
            // res（用户主动关闭验证码）= {ret: 2, ticket: null}
            // res（验证成功） = {ret: 0, ticket: "String", randstr: "String"}
            if (res.ret === 0) {
                $.ajax({
                    dataType: 'json',
                    type: 'POST',
                    url: '{{.UrlPrefix}}/signin',
                    async: 'true',
                    data: {
                        'username': $("#username").val(),
                        'password': $("#password").val(),
                        'token': res.ticket
                    },
                    success: function (data) {
                        location.href = data.data.url
                    },
                    error: function (data) {
                        alert(data.responseJSON.msg);
                    }
                });
            } else {
                alert(data.data.msg);
            }
        }, {});

        {% end %}

        {% if .CaptchaDigits %}

        $(".captcha").on("click",function(){
            location.reload();
        });

        {% end %}

        function submitData() {
            {% if .TencentWaterProofWallData.AppID  %}
            captcha.show();
            {% else %}
            $.ajax({
                dataType: 'json',
                type: 'POST',
                url: '{{.UrlPrefix}}/signin',
                async: 'true',
                data: {
                    'username': $("#username").val(),
                    {% if .CaptchaDigits %}
                        'token': $("#captcha").val() + "," + $("#captcha_id").val(),
                    {% end %}
                    'password': $("#password").val()
                },
                success: function (data) {
                    location.href = data.data.url
                },
                error: function (data) {
                    alert(data.responseJSON.msg);
					location.reload();
                }
            });
            {% end %}
        }
    </script>

    </html>
{{end}}`}

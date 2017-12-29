<div class="header-area">
        <div class="container">
            <div class="row">

                {{if .IsLogin}}
                <div class="col-md-8">
                    <div class="user-menu">
                        <ul>
                            <li><a href="/myaccount"><i class="fa fa-user"></i>{{i18n .Lang "my_account"}}</a></li>
                            <li><a href="#"><i class="fa fa-heart"></i> {{i18n .Lang "my_collect"}}</a></li>
                            <li><a href="cart.html"><i class="fa fa-user"></i> {{i18n .Lang "my_message"}}</a></li>
                            <li><a href="checkout.html"><i class="fa fa-user"></i> {{i18n .Lang "my_bill"}}</a></li>
                                
                        </ul>
                    </div>
                </div>
                
                <div class="col-md-4">
                    <div class="header-right">
                        <ul class="list-unstyled list-inline">
                            <li><a href="/addApply">{{i18n .Lang "personalized_customization"}}</a></li>
                            

                            <li class="dropdown dropdown-small">
                                <a data-toggle="dropdown" data-hover="dropdown" class="dropdown-toggle" href="#"><span class="key">welcome</span><span class="value">{{i18n .Lang "personalized_customization"}} </span><b class="caret"></b></a>
                                <ul class="dropdown-menu">
                                    <li><a href="/logout">{{i18n .Lang "exit"}}</a></li>
                                </ul>
                            </li>
                            <li class="dropdown dropdown-small">
                                <a data-toggle="dropdown" data-hover="dropdown" class="dropdown-toggle" href="#"><span class="key">{{i18n .Lang "language"}}{{.Lang}}:</span><span class="value">{{.username}} </span><b class="caret"></b></a>
                                <ul class="dropdown-menu">
                                    <li><a href="/?lang=zh-CN">{{i18n .Lang "chinese"}}</a></li>
                                    <li><a href="/?lang=en-US">{{i18n .Lang "english"}}</a></li>
                                </ul>
                            </li>
                        </ul>
                    </div>
                </div>

                {{else}}


                <div class="col-md-8">
                    <div class="user-menu">
                        <ul>


                                
                        </ul>
                    </div>
                </div>
                 <div class="col-md-4">
                    <div class="header-right">
                        <ul class="list-unstyled list-inline">
                            <li><a href="/addApply">个性化定制数据</a></li>
                            
                            <li><a href="/login"><i class="fa fa-user"></i>登陆</a></li>
                        </ul>
                    </div>
                </div>
                {{end}}

            </div>
        </div>
    </div> <!-- End header area -->
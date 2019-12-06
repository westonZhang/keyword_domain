package main

import (
	"fmt"
	"keyword_domain/services/other_services"
)

var html = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=gbk" />
    <meta name="renderer" content="webkit">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <title>免费建站_网站建设_免费网站_自助建站-【建站ABC】</title>
    <meta name="keywords" content="免费建站,免费网站,自助建站,网站建设,免费网站申请,免费网站建设,免费建网站,微信网站,手机APP,小程序" />
    <meta name="description" content="建站ABC为中小企业提供免费建站,免费网站,免费微网站,自助建站,网站建设,免费网站申请,免费网站建设，免费建网站,免费微信网站,手机APP,小程序等永久性免费建站服务。" />

        <link href="/static/css/base.css?126=126" rel="stylesheet" type="text/css" />
        <link href="/static/css/css.css?126=121cssac" rel="stylesheet" type="text/css" />
        <link href="/static/css/animate.min.css?126=126" rel="stylesheet" type="text/css" />
        <meta name="baidu-site-verification" content="EPRoqixMTp" />
    <script type="text/javascript" src="/static/js/jquery-1.8.2.min.js"></script>
    <script> window.webCount = '1845075'; </script>
        <script type="text/javascript" src="/static/js/public_js.js?334"></script>
    <script type="text/javascript" src="/static/js/common.js?4"></script>
        
        <script>
            var _hmt = _hmt || [];
            (function() {
                var hm = document.createElement("script");
                //hm.src = "https://hm.baidu.com/hm.js?b2c1a7693fcc44f5b31b08ab60f77d4c";
                hm.src = "https://hm.baidu.com/hm.js?e26f01a6ed945b6d1d435584abf4ebdf";
                var s = document.getElementsByTagName("script")[0];
                s.parentNode.insertBefore(hm, s);
            })();
        </script>
    
    <script>
        window._global = {
            web_id: 'web_id_cbdd56986897ea3eeb62333c9ccb8e61',
            rand: '56164251',
            ip: '221.224.170.90',
            u: '0',
            link_id: '6095e23e187f074940de7b4efe25147a',
            referer: 'http://www.ev123.net',
            customSiteConf: {"1":{"type":"展示板网站委托建站","price":"1500"},"2":{"type":"商城版网站委托建站","price":"3500"},"3":{"type":"尊享版定制委托建站","price":"5000"},"10":{"type":"尊享版定制委托建站","price":"5000"},"11":{"type":"云官网标准版","price":"400","year":{"1":"0","2":"2","3":"3","4":"4","5":"5","6":"6","7":"7","8":"8","9":"9","10":"10"}},"12":{"type":"云官网高级版","price":"1000","year":{"1":"0","2":"2","3":"3","4":"4","5":"5","6":"6","7":"7","8":"8","9":"9","10":"10"}},"13":{"type":"门户标准版","price":"600","year":{"1":"0","2":"2","3":"3","4":"4","5":"5","6":"6","7":"7","8":"8","9":"9","10":"10"}},"14":{"type":"门户高级版","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"},"price":"1200","discount":"0.8","marketPrice":"1500","discountYear":{"1":"0.8","2":"0.5","3":"0.5","4":"0.5","5":"0.5","6":"0.5","7":"0.5","8":"0.5","9":"0.5","10":"0.5"}},"15":{"type":"基础商城版","price":"1298","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"}},"16":{"type":"商城高级版","price":"3200","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"}},"17":{"type":"分销商城版","price":"4600","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"}},"18":{"type":"拼团商城版","price":"2000","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"}},"19":{"type":"拼团商城高级版","price":"2800","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"}},"20":{"type":"批发商城版","price":"2000","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"}},"21":{"type":"批发商城高级版","price":"2599","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"}},"22":{"type":"多用户商城","price":"7800","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"}},"23":{"type":"论坛标准版","price":"700","year":{"1":"0","2":"2","3":"3","4":"4","5":"5","6":"6","7":"7","8":"8","9":"9","10":"10"}},"24":{"type":"论坛高级版","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"},"price":"3040","discount":"0.8","marketPrice":"3800","discountYear":{"1":"0.8","2":"0.5","3":"0.5","4":"0.5","5":"0.5","6":"0.5","7":"0.5","8":"0.5","9":"0.5","10":"0.5"}},"25":{"type":"问答标准版","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"},"price":"1440","discount":"0.8","marketPrice":"1800","discountYear":{"1":"0.8","2":"0.5","3":"0.5","4":"0.5","5":"0.5","6":"0.5","7":"0.5","8":"0.5","9":"0.5","10":"0.5"}},"26":{"type":"问答高级版","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"},"price":"3840","discount":"0.8","marketPrice":"4800","discountYear":{"1":"0.8","2":"0.5","3":"0.5","4":"0.5","5":"0.5","6":"0.5","7":"0.5","8":"0.5","9":"0.5","10":"0.5"}},"27":{"type":"小程序展示版","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"},"price":"944","discount":"0.8","marketPrice":"1180","discountYear":{"1":"0.8"}},"28":{"type":"小程序商城版","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"},"price":"1264","discount":"0.8","marketPrice":"1580","discountYear":{"1":"0.8"}},"30":{"type":"小程序全功能版","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"},"price":"4640","discount":"0.8","marketPrice":"5800","discountYear":{"1":"0.8"}},"29":{"type":"小程序分销版","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"},"price":"1744","discount":"0.8","marketPrice":"2180","discountYear":{"1":"0.8"}},"31":{"type":"拼团商城 标准版（一年8折两年5折）","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"},"price":"4640","discount":"0.8","marketPrice":"5800","discountYear":{"1":"0.8","2":"0.5","3":"0.5","4":"0.5","5":"0.5","6":"0.5","7":"0.5","8":"0.5","9":"0.5","10":"0.5"}},"32":{"type":"视频课程版","price":"2300","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"}},"33":{"type":"小程序多用户版","price":"3744","discount":"0.8","marketPrice":"4680","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"}},"34":{"type":"视频课程版高级版","price":"5800","year":{"1":"0","2":"0","3":"0","4":"0","5":"0","6":"0","7":"0","8":"0","9":"0","10":"0"}}},
            cm: '',
            kd: '',
            activePage: '',
            urlParams: {"site-referer":"www.ev123.net"}
        };
    </script>
    <script src="/static/js/cookies.js"></script>
    <!-- <script src="/static/js/logger.js"></script> -->
</head>
<body>
    <div class="league-alert" id="leagueAlert">
    <img src="/static/images/banner1/league_alert_img.png">
    <b class="close" title="关闭"></b>
    <a class="btn" href="http://www.ev123.net/z.php"></a>
</div>
<div class="league-alert-bg" id="leagueAlertBg"></div>
<div class="index_head">
    <div class="index_head_c">
        <div class="logo"><img src="/static/images/logo.png?1" alt="建站ABC" /></div>
                <div class="nav">
            <a href="/" class="nav_list cur cur1"><i>首页</i><label></label></a>
            <div href="javascript:;" id="drop_down1">
                <i class="drop_down">产品中心</i>
                <label class=""></label>
                <div class="Two_menu" id="Two_menu">
                    <div class="Two_menu_c">
                                    <div class="Two_menu_list">
                                <div class="menu_list_li nth1">
                    <div class="menu_list_li_tit">展示型网站</div>                        <a href="/website_1_14.html">
                            <span><img src="/static/images/Two_menu_img1.png" alt="" /></span>
                            <p>
                                <em>云官网<img src="/images/hot.png"></em>
                                <strong>快速搭建企业官网</strong>
                            </p>
                        </a>                        <a href="/website_1_64.html">
                            <span><img src="/static/images/Two_menu_img2.png" alt="" /></span>
                            <p>
                                <em>政企门户</em>
                                <strong>政府企业的门户型网站</strong>
                            </p>
                        </a>
                </div>                <div class="menu_list_li nth1">
                    <div class="menu_list_li_tit">小程序</div>                        <a href="/website_35_78.html">
                            <span><img src="/static/images/Two_menu_img3.png" alt="" /></span>
                            <p>
                                <em>小程序标准版（四合一）<img src="/images/hot.png"></em>
                                <strong>快速搭建企业官网</strong>
                            </p>
                        </a>                        <a href="/website_35_79.html">
                            <span><img src="/static/images/Two_menu_img4.png" alt="" /></span>
                            <p>
                                <em>小程序分销商城版（四合一）</em>
                                <strong>小程序商城分佣模式</strong>
                            </p>
                        </a>                        <a href="/website_35_80.html">
                            <span><img src="/static/images/Two_menu_img7.png" alt="" /></span>
                            <p>
                                <em>小程序多用户商城版（四合一）</em>
                                <strong>商家入驻，让别人帮你卖货</strong>
                            </p>
                        </a>                        <a href="/website_35_69.html">
                            <span><img src="/static/images/Two_menu_img18.png" alt="" /></span>
                            <p>
                                <em>小程序微圈儿（微信+百度）<img src="/images/hot.png"></em>
                                <strong>打造社群论坛+电商新模式</strong>
                            </p>
                        </a>
                </div>
            </div>            <div class="Two_menu_list">
                                <div class="menu_list_li">
                    <div class="menu_list_li_tit">商城型网站</div>                        <a href="/website_3_18.html" >
                            <span><img src="/static/images/Two_menu_img8.png" alt="" /></span>
                            <p>
                                <em>基础商城（B2C）<img src="/images/hot.png"></em>
                                <strong>小京东直营商城</strong>
                            </p>
                        </a>                        <a href="/website_3_21.html" >
                            <span><img src="/static/images/Two_menu_img7.png" alt="" /></span>
                            <p>
                                <em>分销商城</em>
                                <strong>让别人帮你卖货</strong>
                            </p>
                        </a>                        <a href="/website_3_72.html" >
                            <span><img src="/static/images/Two_menu_img5.png" alt="" /></span>
                            <p>
                                <em>拼团商城</em>
                                <strong>拼团商城，一起拼团吧</strong>
                            </p>
                        </a>                        <a href="/website_3_19.html" >
                            <span><img src="/static/images/Two_menu_img6.png" alt="" /></span>
                            <p>
                                <em>批发商城</em>
                                <strong>批发商城，多买多优惠</strong>
                            </p>
                        </a>                        <a href="/website_3_20.html" >
                            <span><img src="/static/images/Two_menu_img9.png" alt="" /></span>
                            <p>
                                <em>多用户商城（B2B2C）</em>
                                <strong>小淘宝，支持商家入驻</strong>
                            </p>
                        </a>
                </div>                <div class="menu_list_li">
                    <div class="menu_list_li_tit">大数据</div>                        <a href="http://www.ev123.net/d.php" target="_blank">
                            <span><img src="/static/images/zhanqun.png" alt="" /></span>
                            <p>
                                <em>站群系统</em>
                                <strong>为企业打造专属互联网大数据分析中心</strong>
                            </p>
                        </a>
                </div>
            </div>                    <div class="Two_menu_list">
                        <div class="menu_list_li">
                            <div class="menu_list_li_tit">O2O行业网站</div>                        <a href="https://canyin.evyun.cn?cm=" target="_blank">
                            <span><img src="/static/images/Two_menu_img10.png" alt="" /></span>
                            <p>
                                <em>餐饮O2O<img src="/images/hot.png"></em>
                                <strong>外卖堂食预约高效管理</strong>
                            </p>
                        </a>                        <a href="/website_38_31.html" >
                            <span><img src="/static/images/Two_menu_img11.png" alt="" /></span>
                            <p>
                                <em>财税O2O</em>
                                <strong>针对工商财税建造</strong>
                            </p>
                        </a>                        <a href="/website_38_61.html" >
                            <span><img src="/static/images/Two_menu_img13.png" alt="" /></span>
                            <p>
                                <em>家政O2O</em>
                                <strong>家政版块量身定做</strong>
                            </p>
                        </a>                        <a href="http://www.evyun.cn/reg.php?cate=38&type=62&cm=" >
                            <span><img src="/static/images/Two_menu_img14.png" alt="" /></span>
                            <p>
                                <em>培训O2O</em>
                                <strong>针对培训机构打造</strong>
                            </p>
                        </a>                        <a href="http://www.evyun.cn/reg.php?cate=38&type=63&cm=" >
                            <span><img src="/static/images/Two_menu_img15.png" alt="" /></span>
                            <p>
                                <em>酒店O2O</em>
                                <strong>针对自营酒店打造</strong>
                            </p>
                        </a>                        <a href="http://www.evyun.cn/reg.php?cate=38&type=59&cm=" >
                            <span><img src="/static/images/Two_menu_img16.png" alt="" /></span>
                            <p>
                                <em>房产O2O</em>
                                <strong>针对房产中介打造</strong>
                            </p>
                        </a>                        <a href="http://www.evyun.cn/reg.php?cate=38&type=52&cm=" >
                            <span><img src="/static/images/Two_menu_img17.png" alt="" /></span>
                            <p>
                                <em>维修O2O</em>
                                <strong>针对维修行业打造</strong>
                            </p>
                        </a>
                        </div>
                    </div>            <div class="Two_menu_list">
                                <div class="menu_list_li">
                    <div class="menu_list_li_tit">主推</div>                        <a href="http://www.ev123.net/safety.php" >
                            <span><img src="/static/images/protected.png" alt="" /></span>
                            <p>
                                <em><b style='color:red'>安全防护营销型网站</b></em>
                                <strong>安全+一站千词+营销名片</strong>
                            </p>
                        </a>
                </div>                <div class="menu_list_li">
                    <div class="menu_list_li_tit">互动网站</div>                        <a href="/website_7_69.html" >
                            <span><img src="/static/images/Two_menu_img18.png" alt="" /></span>
                            <p>
                                <em>微圈儿<img src="/images/hot.png"></em>
                                <strong>打造社群+电商新模式+同城信息平台</strong>
                            </p>
                        </a>                        <a href="/website_7_15.html" >
                            <span><img src="/static/images/Two_menu_img19.png" alt="" /></span>
                            <p>
                                <em>论坛推广版</em>
                                <strong>传统论坛加手机</strong>
                            </p>
                        </a>                        <a href="/website_7_77.html" >
                            <span><img src="/static/images/Two_menu_img20.png" alt="" /></span>
                            <p>
                                <em>微问答</em>
                                <strong>邀请问答、悬赏问答、付费看回答</strong>
                            </p>
                        </a>                        <a href="/website_7_82.html" >
                            <span><img src="/static/images/Two_menu_img21.png" alt="" /></span>
                            <p>
                                <em>视频课程</em>
                                <strong>视频、音频课程平台解决方案</strong>
                            </p>
                        </a>
                </div>
            </div>
                    </div>
                </div>
            </div>
            <a style="display: none;" href="/website_35_78.html" class="nav_list TopActiveMenu ps " id="xcxTgMenu">
                <i>小程序</i><label></label>
                <div class="Small_pro_div" trigger="menu" style="display:none;">
                    <div class="Small_por_box">
                        <div class="Small_list_a" trigger="bgcolor" style="padding-top:28px;">
                            <a href="/website_35_78.html">
                                <div class="Small_img"><img src="/static/images/Two_menu_img3.png" alt=""></div>
                                <div class="Small_text">
                                    <em>小程序标准版</em>
                                    <span>快速搭建企业官网</span>
                                </div>
                            </a>
                        </div>
                        <div class="Small_list_a" trigger="bgcolor">
                            <a href="/website_35_79.html">
                                <div class="Small_img"><img src="/static/images/Two_menu_img4.png" alt=""></div>
                                <div class="Small_text">
                                    <em>小程序分销商城版</em>
                                    <span>小程序的商城多级分佣</span>
                                </div>
                            </a>
                        </div>
                        <div class="Small_list_a" trigger="bgcolor">
                            <a href="/website_35_80.html">
                                <div class="Small_img"><img src="/static/images/Two_menu_img7.png" alt=""></div>
                                <div class="Small_text">
                                    <em>小程序多用户商城版</em>
                                    <span>让别人帮你卖货</span>
                                </div>
                            </a>
                        </div>
                        <div class="Small_list_a" trigger="bgcolor">
                            <a href="http://wq.ev123.net">
                                <div class="Small_img"><img src="/static/images/Two_menu_img18.png" alt=""></div>
                                <div class="Small_text">
                                    <em>小程序微圈版</em>
                                    <span>打造社群论坛+电商新模式</span>
                                </div>
                            </a>
                        </div>
                    </div>
                </div>
            </a>
                        <a href="https://shop.evyun.cn?site-referer=www.ev123.net" target="_blank" class="nav_list"><i>商城</i><label></label></a>
                        <a href="https://www.hi-mp.com?site-referer=www.ev123.net" target="_blank" class="nav_list"><i>智能名片</i><label></label></a>

            <a href="shop.php" class="nav_list "><i style="color:#f90">商家助手</i><label></label></a>
            <a href="templates.php" class="nav_list "><i>模版</i><label></label></a>
            <a href="custom.php" class="nav_list "><i>定制</i><label></label></a>
            <a href="z.php" class="nav_list "><i>加盟合作</i><label></label></a>
            <a href="http://server.evyun.cn" target="_blank" class="nav_list "><i>服务市场</i><label></label></a>
            <!-- <a href="crm.php" class="nav_list "><i>CRM</i><label></label></a> -->
            <a href="http://admin.evyun.cn/help_index.php?channel_id=12401520" target="_blank" class="nav_list TopActiveMenu ps ">
                <i>帮助中心</i><label></label>
                <div class="Small_pro_div problem_menu" trigger="menu" style="display:none;height:90px">
                    <div class="Small_por_box">
                        <div class="Small_list_a" trigger="bgcolor" style="padding-top:28px;">
                            <a href="http://help.evyun.cn" target="_blank">
                                <div class="Small_img"><img src="/static/images/problem.png" alt=""></div>
                                <div class="Small_text">
                                    <em>问题库</em>
                                    <span>常见问题快速找到答案</span>
                                </div>
                            </a>
                        </div>
                    </div>
                </div>
            </a>
        </div>
        <div class="login">
            <a href="http://admin.evyun.cn/v/home/index.html#/signin?site-referer=www.ev123.net" target="_blank" class="register_a">免费注册</a>
            <a href="http://admin.evyun.cn/login.php?site-referer=www.ev123.net" target="_blank" class="login_a">登录</a>
        </div>
                <dl class="top-welcome" id="topWelcome">
            <dt>
                <span><s>Hi~</s>李先生</span><i class="icon-img icon-1"></i>
            </dt>
            <dd>
                <ul>
                    <li class="cur"><i class="icon-img icon-2"></i><a href="###">用户中心</a></li>
                    <li><i class="icon-img icon-3"></i><a href="###">设计网站</a></li>
                    <li><i class="icon-img icon-4"></i><a href="###">退出登录</a></li>
                </ul>
            </dd>
        </dl>
    </div>
</div>
<script language="javascript">
	if(/Android|webOS|iPhone|iPod|BlackBerry/i.test(navigator.userAgent)) {
		window.location.href = "http://demo.jzabc.cn/dom/wap_site_reg/index.php";
	}
</script>
<div class="index_center">
    <div class="banner" id="banner">
        <div class="banner_img">
            <div class="bg1">
              <div class="bg-con">
                <div class="new-year-txt">
                  <em class="em-1 load-animate" data-animate-name="tada" data-animate-duration="0.50s" data-animate-delay="0.60s">2019，</em>
                  <em class="em-2 load-animate" data-animate-name="fadeInLeft" data-animate-duration="0.50s" data-animate-delay="1.2s">您</em>
                  <em class="em-3 load-animate" data-animate-name="fadeInLeft" data-animate-duration="0.50s" data-animate-delay="1.4s">需</em>
                  <em class="em-4 load-animate" data-animate-name="fadeInLeft" data-animate-duration="0.50s" data-animate-delay="1.6s">要</em>
                  <em class="em-5 load-animate" data-animate-name="fadeInLeft" data-animate-duration="0.50s" data-animate-delay="1.8s">一</em>
                  <em class="em-6 load-animate" data-animate-name="fadeInLeft" data-animate-duration="0.50s" data-animate-delay="2s">个</em>
                  <em class="em-7 load-animate" data-animate-name="fadeInLeft" data-animate-duration="0.50s" data-animate-delay="2.2s">有</em>
                  <em class="em-8 load-animate" data-animate-name="bounceIn" data-animate-duration="1.00s" data-animate-delay="2.6s"></em>
                  <em class="em-9 load-animate" data-animate-name="bounceIn" data-animate-duration="1.00s" data-animate-delay="2.8s"></em>
                  <em class="em-10 load-animate" data-animate-name="fadeInLeft" data-animate-duration="0.50s" data-animate-delay="3.6s">的</em>
                  <em class="em-11 load-animate" data-animate-name="fadeInLeft" data-animate-duration="0.50s" data-animate-delay="3.9s">网</em>
                  <em class="em-12 load-animate" data-animate-name="fadeInLeft" data-animate-duration="0.50s" data-animate-delay="4.2s">站</em>
                </div>
                <div class="banner-fun-list">
                    <b class="load-animate" data-animate-name="bounceIn" data-animate-duration="0.50s" data-animate-delay="4s">电脑站</b>
                    <em class="load-animate" data-animate-name="fadeInDown" data-animate-duration="0.50s" data-animate-delay="4.1s">+</em>
                    <b class="load-animate" data-animate-name="bounceIn" data-animate-duration="0.50s" data-animate-delay="4.2s">手机站</b>
                    <em class="load-animate" data-animate-name="fadeInDown" data-animate-duration="0.50s" data-animate-delay="4.3s">+</em>
                    <b class="load-animate" data-animate-name="bounceIn" data-animate-duration="0.50s" data-animate-delay="4.4s">微信站</b>
                    <em class="load-animate" data-animate-name="fadeInDown" data-animate-duration="0.50s" data-animate-delay="4.5s">+</em>
                    <b class="load-animate" data-animate-name="bounceIn" data-animate-duration="0.50s" data-animate-delay="4.6s">苹果APP</b>
                    <em class="load-animate" data-animate-name="fadeInDown" data-animate-duration="0.50s" data-animate-delay="4.7s">+</em>
                    <b class="load-animate" data-animate-name="bounceIn" data-animate-duration="0.50s" data-animate-delay="4.8s">安卓APP</b>
                    <em class="load-animate" data-animate-name="fadeInDown" data-animate-duration="0.50s" data-animate-delay="4.9s">+</em>
                    <b class="load-animate" data-animate-name="bounceIn" data-animate-duration="0.50s" data-animate-delay="5.0s">小程序</b>
                </div>
                <em class="new-computer load-animate" data-animate-name="fadeInUp" data-animate-duration="0.50s" data-animate-delay="0.00s"></em>
              </div>
            </div>
            <div class="bg2"></div>
            <div class="bg3"></div>
            <div class="bg4"></div>
            <div class="bg5"></div>
            <div class="bg6"></div>
            <div class="bg7"></div>
            <div class="bg8"></div>
            <div class="bg9"></div>
        </div>
        <div class="banner_c">
            <div class="banner_nav load-animate" data-animate-name="zoomIn" data-animate-duration="0.50s" data-animate-delay="0.00s">
                <a href="javascript:;" class="new_one one_a cur">
                    <i class="a-bg"></i>
                    <span>创造能为企业<br />带去订单的网站</span>
                </a>
                <a href="/website_1_14.html" class="new_one one_a2">
                    <i class="a-bg"></i>
                    <em><img src="/static/images/banner_icon1.png" alt="" /></em>
                    <p>展示型</p>
                </a>
                <a href="https://shop.evyun.cn/" class="new_one one_a3" target="_shop">
                    <i class="a-bg"></i>
                    <em><img src="/static/images/banner_icon2.png" alt="" /></em>
                    <p>商城型</p>
                </a>
                <a  href="https://www.xcxabc.net/" target="_xcx"  class="new_one one_a9">
                    <i class="a-bg"></i>
                    <em><img src="/static/images/banner_icon8.png" alt="" /></em>
                    <p>小程序</p>
                </a>
                <a href="https://shop.evyun.cn/s_j.php?cate=3&type=21" class="new_one one_a4" target="_shop">
                    <i class="a-bg"></i>
                    <em><img src="/static/images/banner_icon3.png" alt="" /></em>
                    <p>分　销</p>
                </a>
                <a class="new_one one_a5" href="https://canyin.evyun.cn/" target="_blank">
                    <i class="a-bg"></i>
                    <em><img src="/static/images/banner_icon4.png" alt="" /></em>
                    <p>餐　饮</p>
                </a>
                <a href="http://wq.ev123.net" class="new_one one_a6">
                    <i class="a-bg"></i>
                    <em><img src="/static/images/banner_icon5.png" alt="" /></em>
                    <p>微圈儿</p>
                </a>
                 <a href="/website_7_82.html" class="new_one one_a8">
                    <i class="a-bg"></i>
                    <em><img src="/static/images/banner_icon7.png" alt="" /></em>
                    <p>视频课程</p>
                </a>
                <a href="https://www.hi-mp.com/" target="_blank" class="new_one one_a7">
                    <i class="a-bg"></i>
                    <em><img src="/static/images/banner_icon9.png?123" alt="" /></em>
                    <p>AI智能名片</p>
                </a>
            </div>
            <div class="banner_right">
                <div class="banner_one_newyear">
                    <div class="banner_newyear">
                        <p>六站合一：电脑站</p>
                        <strong>+</strong>
                        <p>手机站</p>
                        <strong>+</strong>
                        <p>微信站</p>
                        <strong>+</strong>
                        <p>苹果APP</p>
                        <strong>+</strong>
                        <p>安卓APP</p>
                        <strong>+</strong>
                        <p>小程序</p>
                    </div>
                    <div class="one_button">
                        <a href="http://www.evyun.cn/reg.php" target="domain" class="register_a1">免费注册</a>
                        <a href="http://admin.evyun.cn/login.php?site-referer=www.ev123.net" target="domain" class="login_a1">管理登录</a>
                    </div>
                    <span><em id="userNumber" data-username="1845075">0</em> 家企业正在使用建站ABC</span>
                </div>
                <div class="banner_two_c">
                    <div class="two_L">
                        <span>设计大气，3分钟搭建上线</span>
                        <p>
                            <em>标配5G空间</em>
                            <em>设计大气</em>
                            <em>专业SEO优化</em>
                            <em>一键生成APP</em>
                        </p>
                        <strong>电脑站+微信+APP+小程序一站式管理<br>维护方便，页面自由拖拽，模块灵活</strong>
                        <div class="two_bottom">
                            <a href="http://www.evyun.cn/reg.php?cate=1&type=14" target="domain" class="taste">免费建站</a>
                            <a href="/website_1_14.html" class="details">查看详情</a>
                        </div>
                    </div>
                    <div class="two_img" data-animate-name="fadeInUp" data-animate-duration="1.00s" data-animate-delay="0.10s"><img src="/static/images/banner_img2.png" alt="" /></div>
                </div>
                <div class="banner_three_c">
                    <div class="three_L">
                        <span>从渠道打通、会员管理，到精准客<br />户营销、全方位覆盖经营场景</span>
                        <p>为你量身打造品牌商城</p>
                        <div class="three_a">
                            <a href="https://shop.evyun.cn/s_j.php?cate=3&type=18" target="_shop" class="three_a1">
                                <strong><img src="/static/images/banner_pic1.png" alt="" /></strong>
                                <em>B2C商城</em>
                            </a>
                            <a href="https://shop.evyun.cn/s_j.php?cate=3&type=19"  target="_shop"class="three_a1">
                                <strong><img src="/static/images/banner_pic2.png" alt="" /></strong>
                                <em>批发商城</em>
                            </a>
                            <a href="https://shop.evyun.cn/s_j.php?cate=3&type=72" target="_shop" class="three_a1">
                                <strong><img src="/static/images/banner_pic3.png" alt="" /></strong>
                                <em>拼团商城</em>
                            </a>
                            <a href="https://shop.evyun.cn/s_j.php?cate=3&type=20" target="_shop" class="three_a2">
                                <strong><img src="/static/images/banner_pic4.png" alt="" /></strong>
                                <em>多用户商城（B2B2C）</em>
                            </a>
                        </div>
                    </div>
                    <div class="three_R" data-animate-name="fadeInUp" data-animate-duration="1.00s" data-animate-delay="0.10s"><img src="/static/images/banner_img3.png" alt="" /></div>
                </div>
                <div class="banner_nine_c">
                    <div class="nine_L">
                        <span>一键接入小程序，拓展更多渠道</span>
                        <p>自由拖拽，无限增加模块</p>
                        <p>展示型、商城型、O2O等多版本</p>
                        <p>全覆盖：微信+百度+支付宝+抖音头条</p>
                        <div class="two_bottom mar_t">
                            <a href="http://www.evyun.cn/reg.php?cate=35&type=78" target="domain" class="taste">免费建站</a>
                            <a href="https://www.xcxabc.net/" target="_xcx" class="details">查看详情</a>
                        </div>
                    </div>
                    <div class="nine_R" data-animate-name="fadeInUp" data-animate-duration="1.00s" data-animate-delay="0.10s"><img src="/static/images/banner_img9.png" alt="" /></div>
                </div>
                <div class="banner_four_c">
                    <div class="four_L">
                        <span>一站式分销解决方案</span>
                        <p>快速建设分销渠道</p>
                        <p>传统分销+万人分销双模式选择</p>
                        <p>分销佣金一键提现</p>
                        <div class="two_bottom mar_t">
                            <a href="http://www.evyun.cn/reg.php?cate=3&type=21" target="domain" class="taste">免费建站</a>
                            <a href="/website_3_21.html" class="details">查看详情</a>
                        </div>
                    </div>
                    <div class="four_R" data-animate-name="zoomIn" data-animate-duration="1.00s" data-animate-delay="0.10s"><img src="/static/images/banner_img4.png" alt="" /></div>
                </div>
                <div class="banner_five_c">
                    <div class="five_L">
                        <span>堂食、外卖、预约一应俱全</span>
                        <p>优惠券、红包、会员折扣价享不停</p>
                        <em>满足顾客各种需求</em>
                        <div class="two_bottom">
                            <a href="http://www.evyun.cn/reg.php?cate=38&type=50" target="domain" class="taste">免费建站</a>
                            <a href="https://canyin.evyun.cn" target="_blank" class="details">查看详情</a>
                        </div>
                    </div>
                    <div class="five_R" data-animate-name="fadeInUp" data-animate-duration="1.00s" data-animate-delay="0.10s"><img src="/static/images/banner_img5.png" alt="" /></div>
                </div>
                <div class="banner_Six_c">
                    <div class="Six_L">
                        <span>为社群经济量身打造</span>
                        <p>集微论坛、社群电商、新媒体运营为一体</p>
                        <em>实现社群+电商统一运营</em>
                        <div class="two_bottom">
                            <a href="http://www.evyun.cn/reg.php?cate=7&type=69" target="domain" class="taste">免费建站</a>
                            <a href="http://wq.ev123.net" class="details">查看详情</a>
                        </div>
                    </div>
                    <div class="Six_R" data-animate-name="fadeInUp" data-animate-duration="1.00s" data-animate-delay="0.10s"><img src="/static/images/banner_img6.png" alt="" /></div>
                </div>
                <div class="banner_eight_c">
                    <div class="eight_L">
                        <span>视频、音频课程平台解决方案</span>
                        <p>适用于培训公司公开课</p>
                        <p>适用于在线教育学习平台</p>
                        <p>适用于视频教程网</p>
                        <div class="two_bottom eight_bottom">
                            <a href="http://www.evyun.cn/reg.php?cate=7&type=82" target="domain" class="taste">免费建站</a>
                            <a href="product.php?cate=7&type=82" class="details">查看详情</a>
                        </div>
                    </div>
                    <div class="eight_R" data-animate-name="fadeInUp" data-animate-duration="1.00s" data-animate-delay="0.10s"><img src="/static/images/banner_img8.png" alt="" /></div>

                </div>
                <!-- <div class="banner_Seven_c">
                    <div class="Seven_L">
                        <span>问答、知识变现解决方案</span>
                        <p>邀请问答、悬赏问答、付费看回答<br/>专家库、专家一对一提问</p>
                        <div class="two_bottom">
                            <a href="http://www.evyun.cn/reg.php?cate=7&type=77" target="domain" class="taste">免费建站</a>
                            <a href="product.php?cate=7&type=77" class="details">查看详情</a>
                        </div>
                    </div>
                    <div class="Seven_R" data-animate-name="fadeInUp" data-animate-duration="1.00s" data-animate-delay="0.10s"><img src="/static/images/banner_img7.png" alt="" /></div>
                </div> -->
                <div class="banner_Seven_c">
                    <div class="Seven_L">
                        <span>智能名片，连接微信10亿用户</span>
                        <p>打造个人专属形象</p>
                        <p>线上线下分享，激活人脉迅速成交</p>
                        <p>无纸化社交传播，用时短，效率高</p>
                        <p>不加好友，随时沟通</p>
                        <div class="two_bottom">
                            <a href="http://www.evyun.cn/reg.php?cate=104&type=71" target="domain" class="taste">免费注册</a>
                            <a href="https://www.hi-mp.com/ " target="_blank" class="details">查看详情</a>
                        </div>
                    </div>
                    <div class="Seven_R" data-animate-name="fadeInUp" data-animate-duration="1.00s" data-animate-delay="0.10s"><img src="/static/images/banner_img10.png" alt="" /></div>
                </div>


            </div>
        </div>
    </div>
    <div class="advantage">
        <div class="advantage_c">
            <div class="advantage_tit">建站免费、快速、灵活</div>
            <div class="advantage_con">
                <div class="advantage_list load-animate" data-animate-name="fadeInUp" data-animate-duration="0.50s" data-animate-delay="0.50s">
                    <span><img src="/static/images/advantage_img1.png" alt="" /></span>
                    <p>千种精品网站模板</p>
                    <em>上百个行业，设计师精心打磨，<br />看上哪款选哪款！</em>
                </div>
                <div class="advantage_list load-animate" data-animate-name="fadeInUp" data-animate-duration="0.50s" data-animate-delay="1.00s">
                    <span><img src="/static/images/advantage_img2.png" alt="" /></span>
                    <p>自由拖拽模块</p>
                    <em>一键拖拽添加模块，摆放位<br />置随心拽！</em>
                </div>
                <div class="advantage_list load-animate" data-animate-name="fadeInUp" data-animate-duration="0.50s" data-animate-delay="1.50s">
                    <span><img src="/static/images/advantage_img3.png" alt="" /></span>
                    <p>三分钟极速建站</p>
                    <em>点击建站即刻生成，无需技术<br />也能随心建站！</em>
                </div>
                <div class="advantage_list load-animate" data-animate-name="fadeInUp" data-animate-duration="0.50s" data-animate-delay="2.00s">
                    <span><img src="/static/images/advantage_img4.png" alt="" /></span>
                    <p>丰富的素材库</p>
                    <em>无需找图修图，图标美图等<br />尽在素材库！</em>
                </div>
            </div>
            <div class="page-register-btn"><a target="_blank" href="http://admin.evyun.cn/v/home/index.html#/signin?site-referer=www.ev123.net">免费注册</a></div>
        </div>
    </div>
    <div class="system">
        <div class="system_tit">功能全、营销体系完整</div>
        <p class="system_titmin">20款商城营销功能、41款应用插件仍在不断更新</p>
        <div class="click_a">
            <a href="javascript:;" class="hover_L"><img src="/static/images/hover_L.png" alt="" /></a>
            <a href="javascript:;" class="hover_R"><img src="/static/images/hover_R.png" alt="" /></a>
        </div>
        <div class="system_con">
            <div class="click_con">
                <div class="system_list">
                    <div class="system_list_tit">
                        <span><em style=" background:#ff6666;"></em><strong>会员</strong></span>
                        <p>成为会员吧，优惠多多</p>
                    </div>
                    <div class="system_list_con">
                        <div class="min_list">
                            <span><img src="/static/images/system_img1.jpg" alt="" /></span>
                            <p>
                                <em>会员管理</em>
                                <strong>管理会员信息</strong>
                            </p>
                        </div>
                        <div class="min_list">
                            <span><img src="/static/images/system_img2.jpg" alt="" /></span>
                            <p>
                                <em>会员等级</em>
                                <strong>会员等级不同享受不同折扣</strong>
                            </p>
                        </div>
                        <div class="min_list">
                            <span><img src="/static/images/system_img3.jpg" alt="" /></span>
                            <p>
                                <em>会员特权</em>
                                <strong>享受会员给您带来的特权吧</strong>
                            </p>
                        </div>
                    </div>
                    <div class="Mask"></div>
                </div>
                <div class="system_list">
                    <div class="system_list_tit">
                        <span><em style=" background:#d74aff;"></em><strong>营销</strong></span>
                        <p>多样营销，带来流量与销量</p>
                    </div>
                    <div class="system_list_con">
                        <div class="min_list">
                            <span><img src="/static/images/system_img4.jpg" alt="" /></span>
                            <p>
                                <em>优惠券</em>
                                <strong>优惠活动享不停</strong>
                            </p>
                        </div>
                        <div class="min_list">
                            <span><img src="/static/images/system_img5.jpg" alt="" /></span>
                            <p>
                                <em>红包</em>
                                <strong>简单粗暴让粉丝来的更猛烈</strong>
                            </p>
                        </div>
                        <div class="min_list">
                            <span><img src="/static/images/system_img6.jpg" alt="" /></span>
                            <p>
                                <em>积分商城</em>
                                <strong>积分也能付款，积分也能兑换商品</strong>
                            </p>
                        </div>
                    </div>
                    <div class="Mask"></div>
                </div>
                <div class="system_list new_cur">
                    <div class="system_list_tit">
                        <span><em></em><strong>促销</strong></span>
                        <p>多样营销，带来流量与销量</p>
                    </div>
                    <div class="system_list_con">
                        <div class="min_list">
                            <span><img src="/static/images/system_img7.jpg" alt="" /></span>
                            <p>
                                <em>会员卡</em>
                                <strong>会员卡可享受更高折扣</strong>
                            </p>
                        </div>
                        <div class="min_list">
                            <span><img src="/static/images/system_img8.jpg" alt="" /></span>
                            <p>
                                <em>支付卡</em>
                                <strong>各种支付卡折扣享不停</strong>
                            </p>
                        </div>
                        <div class="min_list">
                            <span><img src="/static/images/system_img9.jpg" alt="" /></span>
                            <p>
                                <em>促销海报</em>
                                <strong>炫酷海报让你更加醒目</strong>
                            </p>
                        </div>
                    </div>
                    <div class="Mask"></div>
                </div>
                <div class="system_list">
                    <div class="system_list_tit">
                        <span><em style="background:#2fb3ff;"></em><strong>互动</strong></span>
                        <p>随心互动，让你更了解客户</p>
                    </div>
                    <div class="system_list_con">
                        <div class="min_list">
                            <span><img src="/static/images/system_img10.jpg" alt="" /></span>
                            <p>
                                <em>微邀请</em>
                                <strong>颠覆传统方式，让请帖更快捷</strong>
                            </p>
                        </div>
                        <div class="min_list">
                            <span><img src="/static/images/system_img11.jpg" alt="" /></span>
                            <p>
                                <em>微投票</em>
                                <strong>了解粉丝选择，让你的特色更鲜明</strong>
                            </p>
                        </div>
                        <div class="min_list">
                            <span><img src="/static/images/system_img12.jpg" alt="" /></span>
                            <p>
                                <em>微留言</em>
                                <strong>用户的互动交流利器</strong>
                            </p>
                        </div>
                    </div>
                    <div class="Mask"></div>
                </div>
                <div class="system_list">
                    <div class="system_list_tit">
                        <span><em></em><strong>商城 </strong></span>
                        <p>商城组件功能完善，满足你的各种需求</p>
                    </div>
                    <div class="system_list_con">
                        <div class="min_list">
                            <span><img src="/static/images/system_img13.jpg" alt="" /></span>
                            <p>
                                <em>一元购</em>
                                <strong>一元秒杀活动来袭</strong>
                            </p>
                        </div>
                        <div class="min_list">
                            <span><img src="/static/images/system_img14.jpg" alt="" /></span>
                            <p>
                                <em>拼团购</em>
                                <strong>拼团折扣价，价钱低到爆</strong>
                            </p>
                        </div>
                        <div class="min_list">
                            <span><img src="/static/images/system_img15.jpg" alt="" /></span>
                            <p>
                                <em>满减</em>
                                <strong>满减优惠来啦</strong>
                            </p>
                        </div>
                    </div>
                    <div class="Mask"></div>
                </div>
                <div class="system_list">
                    <div class="system_list_tit">
                        <span><em style="background:#cc53ff"></em><strong>游戏活动</strong></span>
                        <p>游戏活动，边玩边赚红包</p>
                    </div>
                    <div class="system_list_con">
                        <div class="min_list">
                            <span><img src="/static/images/system_img16.jpg" alt="" /></span>
                            <p>
                                <em>摇一摇</em>
                                <strong>通过摇一摇领取红包</strong>
                            </p>
                        </div>
                        <div class="min_list">
                            <span><img src="/static/images/system_img17.jpg" alt="" /></span>
                            <p>
                                <em>刮刮卡</em>
                                <strong>刮刮卡刮出折扣</strong>
                            </p>
                        </div>
                        <div class="min_list">
                            <span><img src="/static/images/system_img18.jpg" alt="" /></span>
                            <p>
                                <em>大转盘</em>
                                <strong>转到几折享几折</strong>
                            </p>
                        </div>
                    </div>
                    <div class="Mask"></div>
                </div>
            </div>
        </div>
    </div>
    <div class="major">
        <div class="major_c">
            <div class="major_tit">专业的服务市场+专家讲座</div>
            <div class="min_major_tit"><em>为您省心</em></div>
            <div class="major_c1">
                <div class="major_L">
                    <a href="javascript:;" class="major_list cur" onmouseover="message(this,1)">
                        <span><img src="/static/images/major_img1.png" alt="" /></span>
                        <p>帮您设计</p>
                    </a>
                    <a href="javascript:;" class="major_list" onmouseover="message(this,2)">
                        <span><img src="/static/images/major_img2.png" alt="" /></span>
                        <p>帮您对接</p>
                    </a>
                    <a href="javascript:;" class="major_list" onmouseover="message(this,3)">
                        <span><img src="/static/images/major_img3.png" alt="" /></span>
                        <p>帮您上架</p>
                    </a>
                </div>
                <div class="min_major_c">
                    <div class="list1" style="display:block">
                        <div class="list1_bg"><img src="/static/images/major_pic1.jpg" alt="" /></div>
                        <div class="list1_c_bg">
                            <div class="comma1"><img src="/static/images/comma1.png" alt="" /></div>
                            <div class="comma_con">
                                <div class="comma_con_tit">
                                    <span><img src="/static/images/comma_img1.png" alt="" /></span>
                                    <p>帮您设计</p>
                                    <!--<a href="javascript:;" class="link_detail">了解详情</a>-->
                                </div>
                                <strong>无需您自己操作，专业设计帮您量身定做，打造您的专属网站，让您的网站不再和人“撞衫”。</strong>
                            </div>
                            <div class="comma2"><img src="/static/images/comma2.png" alt="" /></div>
                        </div>
                    </div>
                    <div class="list1">
                        <div class="list1_bg"><img src="/static/images/major_pic2.jpg" alt="" /></div>
                        <div class="list1_c_bg">
                            <div class="comma1"><img src="/static/images/comma1.png" alt="" /></div>
                            <div class="comma_con">
                                <div class="comma_con_tit">
                                    <span><img src="/static/images/comma_img2.png" alt="" /></span>
                                    <p>帮您对接</p>
                                    <!--<a href="javascript:;" class="link_detail">了解详情</a>-->
                                </div>
                                <strong>第三方支付、登录等对接帮您搞定，再也不用烦恼第三方对接的问题啦！专人为您调试接口，VIP待遇任您享受。</strong>
                            </div>
                            <div class="comma2"><img src="/static/images/comma2.png" alt="" /></div>
                        </div>
                    </div>
                    <div class="list1">
                        <div class="list1_bg"><img src="/static/images/major_pic3.jpg" alt="" /></div>
                        <div class="list1_c_bg">
                            <div class="comma1"><img src="/static/images/comma1.png" alt="" /></div>
                            <div class="comma_con">
                                <div class="comma_con_tit">
                                    <span><img src="/static/images/comma_img3.png" alt="" /></span>
                                    <p>帮您上架</p>
                                    <!--<a href="javascript:;" class="link_detail">了解详情</a>-->
                                </div>
                                <strong>从一键生成APP到上架，一条龙服务，告别繁杂的操作和长时间的等待流程。来建站ABC，我们帮您上架。</strong>
                            </div>
                            <div class="comma2"><img src="/static/images/comma2.png" alt="" /></div>
                        </div>
                    </div>
                    <div class="list1">
                        <div class="list1_bg"><img src="/static/images/major_pic4.jpg" alt="" /></div>
                        <div class="list1_c_bg">
                            <div class="comma1"><img src="/static/images/comma1.png" alt="" /></div>
                            <div class="comma_con">
                                <div class="comma_con_tit">
                                    <span><img src="/static/images/comma_img4.png" alt="" /></span>
                                    <p>帮您建站</p>
                                   <!-- <a href="javascript:;" class="link_detail">了解详情</a>-->
                                </div>
                                <strong>不会技术也无妨，专业建站帮您一键搭建网站，不再为技术担忧，来建站ABC，享受专业的建站服务。</strong>
                            </div>
                            <div class="comma2"><img src="/static/images/comma2.png" alt="" /></div>
                        </div>
                    </div>
                    <div class="list1">
                        <div class="list1_bg"><img src="/static/images/major_pic5.jpg" alt="" /></div>
                        <div class="list1_c_bg">
                            <div class="comma1"><img src="/static/images/comma1.png" alt="" /></div>
                            <div class="comma_con">
                                <div class="comma_con_tit">
                                    <span><img src="/static/images/comma_img5.png" alt="" /></span>
                                    <p>帮您维护</p>
                                    <!--<a href="javascript:;" class="link_detail">了解详情</a>-->
                                </div>
                                <strong>帮你上传维护内容，无需复杂操作，一切内容我们来帮您解决，省去无数烦恼。</strong>
                            </div>
                            <div class="comma2"><img src="/static/images/comma2.png" alt="" /></div>
                        </div>
                    </div>
                    <div class="list1">
                        <div class="list1_bg"><img src="/static/images/major_pic6.jpg" alt="" /></div>
                        <div class="list1_c_bg">
                            <div class="comma1"><img src="/static/images/comma1.png" alt="" /></div>
                            <div class="comma_con">
                                <div class="comma_con_tit">
                                    <span><img src="/static/images/comma_img6.png" alt="" /></span>
                                    <p>学院培训</p>
                                   <!-- <a href="javascript:;" class="link_detail">了解详情</a>-->
                                </div>
                                <strong>专家培训，为您量身定做培训方案，随时随地帮您排忧解难，答疑解惑。</strong>
                            </div>
                            <div class="comma2"><img src="/static/images/comma2.png" alt="" /></div>
                        </div>
                    </div>
                </div>
                <div class="major_L">
                    <a href="javascript:;" class="major_list" onmouseover="message(this,4)">
                        <span><img src="/static/images/major_img4.png" alt="" /></span>
                        <p>帮您建站</p>
                    </a>
                    <a href="javascript:;" class="major_list" onmouseover="message(this,5)">
                        <span><img src="/static/images/major_img5.png" alt="" /></span>
                        <p>帮您维护</p>
                    </a>
                    <a href="javascript:;" class="major_list" onmouseover="message(this,6)">
                        <span><img src="/static/images/major_img6.png" alt="" /></span>
                        <p>学院培训</p>
                    </a>
                </div>
            </div>
            <div class="page-register-btn"><a target="_blank" href="http://admin.evyun.cn/v/home/index.html#/signin?site-referer=www.ev123.net">免费注册</a></div>
        </div>
    </div>
    <div class="templates-row">
        <div class="row-zz"></div>
        <div class="templates-row-c">
            <h5>海量精美模板，让您与众不同</h5>
            <p>我们有1000+无限绝佳的选项，绝对不重复</p>
            <div class="btn-area">
                <a href="http://admin.evyun.cn/v/home/index.html#/signin?site-referer=www.ev123.net" target="_blank" class="btn">免费注册使用</a>
            </div>
        </div>
    </div>
    <div class="industry">
        <div class="industry_c">
            <div class="industry_tit">互联网+新行业</div>
            <div class="industry_con">
                <div class="industry_list load-animate" data-animate-name="zoomIn" data-animate-duration="1.00s" data-animate-delay="0.20s">
                    <span><a href="https://canyin.evyun.cn/" target="_blank"><img src="/static/images/industry_img1.png" alt="餐饮免费建站" /></a></span>
                    <p>餐饮专版</p>
                </div>
                <div class="industry_list load-animate" data-animate-name="zoomIn" data-animate-duration="1.00s" data-animate-delay="0.40s">
                    <span><a href="http://www.ev123.net/website_38_31.html" ><img src="/static/images/industry_img2.png" alt="财税免费建站" /></a></span>
                    <p>财税专版</p>
                </div>
                <div class="industry_list load-animate" data-animate-name="zoomIn" data-animate-duration="1.00s" data-animate-delay="0.60s">
                    <span><a href="http://www.ev123.net/website_38_61.html" ><img src="/static/images/industry_img3.png"  alt="家政免费建站" /></a></span>
                    <p>家政专版</p>
                </div>
                <div class="industry_list load-animate" data-animate-name="zoomIn" data-animate-duration="1.00s" data-animate-delay="0.80s">
                    <span><a href="http://www.evyun.cn/reg.php?cate=38&type=62" target="_blank"><img src="/static/images/industry_img4.png" alt="培训免费建站" /></a></span>
                    <p>培训专版</p>
                </div>
                <div class="industry_list load-animate" data-animate-name="zoomIn" data-animate-duration="1.00s" data-animate-delay="1.00s">
                    <span><a href="/reg.php?cate=38&type=63" target="_blank"><img src="/static/images/industry_img5.png" alt="酒店免费建站" /></a></span>
                    <p>酒店专版</p>
                </div>
                <div class="industry_list load-animate" data-animate-name="zoomIn" data-animate-duration="1.00s" data-animate-delay="1.20s">
                    <span><a href="/templates.php" target="_blank"><img src="/static/images/industry_img6.png" alt="" /></a></span>
                    <p>&nbsp;</p>
                </div>
            </div>
            <div class="page-register-btn"><a target="_blank" href="http://admin.evyun.cn/v/home/index.html#/signin?site-referer=www.ev123.net">免费注册</a></div>
        </div>
    </div>
    <div class="product-number-info" id="productNumber">
        <div class="product-number-inner">
            <ul class="q">
                <li class="li-1"><b><i id="productNumber_1" data-sum="1845075">0</i><sup>+</sup></b><p>公司优选</p></li>
                <li class="li-2"><b><i id="productNumber_2" data-sum="3870">0</i><em>天</em></b><p>安全运行</p></li>
                <li class="li-3"><b><i id="productNumber_3" data-sum="150">0</i><sup>+</sup></b><p>覆盖行业和类目</p></li>
                <li class="li-4"><b><i id="productNumber_4" data-sum="1663">0</i><em>万</em></b><p>件商品销售中</p></li>
                <li class="li-5"><b><i id="productNumber_5" data-sum="4330119">0</i><sup>+</sup></b><p>个订单为企业带来</p></li>
            </ul>
        </div>
    </div>
    <div class="copyright">
    <div class="copyright_c">
        <div class="copyright_head">
            <div class="copyright_L">
                <p><img src="/static/images/logo.png" alt="" /></p>
                <span>
						<em>服务时间：</em>
						<em>周一至周日  9:00-18:00</em>
					</span>
                <span>
						<em>客服电话：</em>
						<strong>4000-123-127</strong>
					</span>
                <span style="display: none;">
                    <em>客服QQ：</em>
                                        <a href="http://crm2.qq.com/page/portalpage/wpa.php?uin=4000123127&f=1&ty=1&aty=0&a=&from=5" target="_blank">在线交谈</a>
                                    </span>
            </div>
            <div class="copyright_C" >
                <div class="copyright_C_list">
                    <p>关于我们</p>
                    <div class="about">
                        <a href="about_us.php">公司简介</a>
                    </div>
                    <div class="about">
                        <a href="about_us.php#tag1">知识产权</a>
                    </div>
                    <div class="about">
                        <a href="about_us.php#tag2">联系我们</a>
                    </div>
                </div>
                <div class="copyright_C_list">
                    <p>产品服务</p>
                    <div class="about">
                        <a href="/website_1_14.html">云官网</a>
                        <a href="/website_3_20.html">云商城</a>
                    </div>
                    <div class="about">
                        <a href="/website_3_21.html">云分销</a>
                        <a href="http://wq.ev123.net">微圈儿</a>
                    </div>
                    <div class="about">
                        <a href="http://www.xcxabc.net">小程序</a>
                        <a href="/t.php">小程序加盟</a>
                    </div>
                </div>
                <div class="copyright_C_list">
                    <p>建站学院</p>
                    <div class="about">
                        <a href="http://server.evyun.cn/" target="_blank">服务市场</a>
                        <a href="/city.php" target="_blank">城市建站</a>
                    </div>
                    <div class="about">
                        <a href="http://admin.evyun.cn/help_index.php?channel_id=12401520" target="_blank">帮助中心</a>
                         <a href="https://m.kuaidi100.com/" target="_blank">快递查询</a>
                    </div>
                    <div class="about">
                        <a href="http://xy.ev123.net/" target="_blank">建站学院</a>
                    </div>
                </div>
            </div>
            <div class="copyright_R">
                                <span><img src="/static/images/ewm1.jpg?1" alt="" /></span>
                                <p>扫描关注微信公众号</p>
            </div>
        </div>

                <div class="copyright_bottom">
            <span>京B2-20160081 京公海网安备11010802020283号 京ICP备10008488号-32   </span>
            <p>建站ABC创建能为您带去订单的网站</p>
        </div>
    </div>
</div>
</div>
<div class="suspend" id="onLineServe">
    <a href="javascript:;" data-href="https://www.sobot.com/chat/pc_new/index.html?sysNum=31dbf2f454f84fbeb14b92d08a5a6357&channelFlag=2&partnerId=1575530356" data-target="_blank" class="suspend_R">
        <span class="serve-icon"><img src="/static/images/Smile.png" /></span>
        <p>在线客服</p>
    </a>
    <div class="suspend_C">
        <div class="suspend_list">
            <div class="suspend_tit">
                <em>直销客服</em>
            </div>
            <p>4000-123-127</p>
            <span style="text-align: center;width: 100%;display: block;line-height: 22px;">为保证服务质量，请先登录</span>
            <!--<a style="width: 84px;background: url(/static/images/s_btn.png) no-repeat 13px center;" href="https://www.sobot.com/chat/pc_new/index.html?sysNum=31dbf2f454f84fbeb14b92d08a5a6357&channelFlag=2&partnerId=1575530356" target="_blank"></a>-->
            <a style="width: 84px;background: url(/static/images/s_btn.png) no-repeat 13px center;" href="http://admin.evyun.cn" target="_blank"></a>
        </div>
        <div class="suspend_list suspend_list1">
            <div class="suspend_tit">
                <em>代理加盟</em>
            </div>
            <p>4009-990-465</p>
            <span style="text-align: center;width: 100%;display: block;line-height: 22px;">为保证服务质量，请先登录</span>
            <!--<a style="width: 84px;background: url(/static/images/s_btn.png) no-repeat 13px center;" href="https://www.sobot.com/chat/pc_new/index.html?sysNum=31dbf2f454f84fbeb14b92d08a5a6357&partnerId=1575530356" target="_blank"></a>-->
            <a style="width: 84px;background: url(/static/images/s_btn.png) no-repeat 13px center;" href="http://dls.dlszywz.cn/login.php" target="_blank"></a>
        </div>
        <div class="suspend_ewm" style="    padding-top: 20px;">
                        <span><img src="/static/images/ewm1.jpg?1" alt="" /></span>
                        <p>扫一扫，关注我们</p>
        </div>
    </div>
</div><!---->
<img src="/images/stat.png?username=nicky499&url=www.ev123.net&r=1575530356&session_id=ss3hctgjv8o1g2q6l1oujq3pk6&source_visitors_seo=1&proxy_source_else=1&sem_else=1&seo_else=1&computer=1" width="1" height="1">
</body>
</html>
`

func main() {
	keywords, err := other_services.ParseKeywords(html)
	if err != nil {
		panic(err)
	}

	fmt.Println(keywords)
}

//func main() {
//
//}
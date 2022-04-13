
function getPoster() {
    var pos=localStorage.getItem("posId");
    var Url = window.location.pathname;
    var priK=prompt("please input your key:","share key")
    Url = Url.substring(Url.lastIndexOf("/")+1,Url.length);
    $.ajax({
        type: 'POST',
        url: "/getPoster",
        dataType: "JSON",
        data:{
            url:Url,
            priK:priK,
        },
        success: function (data) {
            content = data.content
            $("#getContent").append(content)
            $("#getPoster").append(pos)
            console.log(data.content)
            hljs.initHighlightingOnLoad()
        }
    });

}
function search() {
    $.ajax({
        type:'POST',
        url:"/search",
        dataType:'JSON',
        data:{
            name:$('search').val(),
        },
    });
    
}
function submit() {
    selectLang();
    var selecSyn = localStorage.getItem("syntaxSelected")
    $.ajax({
        type:"POST",
        url:"/submit",
        data:{
            poster:$("#poster").val(),
            syntax: selecSyn,
            content:$("#content").val(),
        },
    });
    // alert("调用成功")
    localStorage.setItem("posId",$("#poster").val())

}
function getUrl() {
    $.ajax({
        type:'POST',
        url:'/getUrl',
        dataType: 'JSON',
        success: function (data) {
            str  = data.url
            console.log("getUrl ",data.url)
            exhibit = "/" + data.url
            localStorage.setItem("exhibit",exhibit)
            $("#pasteHref").attr("href", exhibit)
            localStorage.setItem("url",str)
            Insert()
            console.log("本地保存 ",localStorage.getItem("url"))
        }
    });
}
function copy(text){
    var textareaC = document.createElement('textarea');
    textareaC.setAttribute('readonly', 'readonly'); //设置只读属性防止手机上弹出软键盘
    textareaC.value = text;
    document.body.appendChild(textareaC); //将textarea添加为body子元素
    textareaC.select();
    var res = document.execCommand('copy');
    document.body.removeChild(textareaC);//移除DOM元素
    console.log("复制成功");
    return res;
}
function Insert(){
    // getUrl()
    var nextUrl =localStorage.getItem("url")
    console.log("查找 ",localStorage.getItem("url"))
    $.ajax({
        type:'POST',
        url:'/urlBind',
        dataType: 'JSON',
        data:{
          url:nextUrl,
        },
        success: function (data) {
            key = data.key
            var r=confirm(key)
            if(r==true)
            {
                copy(key)
                alert("复制成功")
            }
        },

    });
}
function selectLang() {
    var myselect=document.getElementById("select")
    var index=myselect.selectedIndex;
    localStorage.setItem("syntaxSelected",myselect.options[index].text)
    console.log(myselect.options[index].text)
}
function appendLang() {
    var selectSyn = localStorage.getItem("syntaxSelected")
    var cname = "language-" + selectSyn
    $("#getContent").attr("class",cname)
}
function appendoption() {
    $("#select").append(
        " <option  disabled>select language</option>"+
        " <option  >bash</option>"+
        " <option >css</option>"+
    "   <option >c++</option>" +
        "  <option  >c#</option>" +
        "  <option  >diff</option>" +
        "  <option  >markdown</option>" +
        "<option >html</option>" +
        "<option >go</option>" +
        "<option >java</option>" +
        "<option >javascript</option>" +
        "<option >json</option>" +
        "<option >lua</option>" +
        "<option >obj-c</option>" +
        "<option >php</option>" +
        "<option >python</option>"+
        "<option >r</option>"+
        "<option >ruby</option>"+
        "<option >rust</option>"+
        "<option >scss</option>"+
        "<option >sql</option>"+
        "<option >vb</option>"+
        "<option >xml</option>"

    )

}
// function Confirm() {
//     var input = prompt("请输入您的密钥",'share key')
//     $.ajax({
//         type:"POST",
//         url:"/confirm"
//     });
//
// }

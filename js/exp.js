
function getPoster() {
    var pos=localStorage.getItem("posId");
    // alert(pos);
    var url = window.location.pathname;
    url = url.substring(url.lastIndexOf("/")+1,url.length);
    // alert(url);
    $.ajax({
        type: 'POST',
        url: "/getPoster",
        dataType: "JSON",
        data:{
            url:url,
        },
        success: function (data) {

            content = data.content
            $("#getContent").append(content)
            $("#getPoster").append(pos)
            console.log(data.content)
            // alert(url)
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
    localStorage.setItem("posId",$("#poster").val())

}
function getUrl() {
    $.ajax({
        type:'POST',
        url:'/getUrl',
        dataType: 'JSON',
        success: function (data) {
            str  = data.url
            exhibit = "/" + data.url
            // alert(exhibit)
            localStorage.setItem("exhibit",exhibit)
            $("#pasteHref").attr("href", exhibit)
            // windows.location.href=exhibit;
            localStorage.setItem("url",str)
        }
    });
}

function Insert(){
    var nextUrl =localStorage.getItem("url")
    // alert(nextUrl)
    $.ajax({
        type:'POST',
        url:'/urlBind',
        dataType: 'JSON',
        data:{
          url:nextUrl,
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
    // class="language-"
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

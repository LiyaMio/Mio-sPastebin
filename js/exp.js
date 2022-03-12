
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
            alert(data.content)
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
    $.ajax({
        type:"POST",
        url:"/submit",
        data:{
            poster:$("#poster").val(),
            syntax:$("#syntax").val(),
            content:$("#content").val(),
        },
    });

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
            $("#pasteHref").attr("href", exhibit)
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
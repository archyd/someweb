function openNav() {

  let sidebarState = document.getElementById("openbtn").innerText;
  console.log(sidebarState)

  if (sidebarState === ">>"){
    document.getElementById("mySidebar").style.width = "250px";
    document.getElementById("main").style.marginLeft = "250px";
    document.getElementById("openbtn").style.marginLeft = "250px";
    document.getElementById("openbtn").innerHTML = "<<";
  }
  else{
    document.getElementById("mySidebar").style.width = "0";
    document.getElementById("main").style.marginLeft= "0";
    document.getElementById("openbtn").style.marginLeft = "0";
    document.getElementById("openbtn").innerHTML = ">>";
  }


    // document.getElementById("mySidebar").style.width = "250px";
    // document.getElementById("main").style.marginLeft = "250px";
    // document.getElementById("openbtn").style.marginLeft = "250px";
    // document.getElementById("openbtn").innerHTML = "<<";
    
  }
  
  function closeNav() {
    document.getElementById("mySidebar").style.width = "0";
    document.getElementById("main").style.marginLeft= "0";
    
    
    document.getElementById("openbtn").style.marginLeft = "0";
    document.getElementById("openbtn").innerHTML = ">>";
  }
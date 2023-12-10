(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[276],{6413:function(e,t,s){(window.__NEXT_P=window.__NEXT_P||[]).push(["/client/instance",function(){return s(8970)}])},9021:function(e,t,s){"use strict";var i=s(5893);s(7294);var r=s(948),n=s(629),o=s(3946),a=s(1703),l=s(44),d=s(5675),x=s.n(d),c=s(1664),p=s.n(c);let h=(0,r.ZP)(n.Z)({position:"sticky",zIndex:"100",display:"flex",backgroundColor:"#FFFFFF",fontSize:"20px",fontFamily:"'Montserrat'",alignItems:"center",gap:"10px",paddingLeft:"10px",marginBottom:"10px",width:"100%",height:"75px",top:"0",boxShadow:"0px 0px 5px 0px #D1D1D1",borderRadius:"10px",boxSizing:"border-box"});(0,r.ZP)(o.Z)({marginRight:"10px"}),t.Z=e=>{let{currentTab:t,onTabChange:s}=e;return(0,i.jsxs)(h,{elevation:1,children:[(0,i.jsx)(p(),{href:"/client/dashboard",onClick:()=>{s(void 0,"Dashboard")},children:(0,i.jsx)(x(),{src:"/assets/logo.jpg",alt:"Company Logo",width:"150",height:"40",borderRadius:"50%"})}),(0,i.jsx)("p",{children:"v0.0.13"}),["Dashboard","Chat","Apps"].includes(t)?(0,i.jsxs)(a.Z,{value:t,onChange:s,sx:{flexGrow:1},children:[(0,i.jsx)(l.Z,{label:"Dashboard",value:"Dashboard"}),(0,i.jsx)(l.Z,{label:"Chat",value:"Chat"}),(0,i.jsx)(l.Z,{label:"Apps",value:"Apps"})]}):(0,i.jsxs)(a.Z,{value:t,onChange:s,sx:{flexGrow:1},children:[(0,i.jsx)(l.Z,{component:"a",href:"/client/dashboard",label:"Dashboard",value:"Dashboard"}),(0,i.jsx)(l.Z,{component:"a",href:"/client/dashboard?tab=Chat",label:"Chat",value:"Chat"}),(0,i.jsx)(l.Z,{component:"a",href:"/client/dashboard?tab=Apps",label:"Apps",value:"Apps"})]})]})}},8970:function(e,t,s){"use strict";s.r(t),s.d(t,{default:function(){return S}});var i=s(5893),r=s(7294),n=s(9417),o=s(5616),a=s(3791);s(7885);let l=[{start:new Date(2020,1,1),end:new Date(2020,1,2),name:"Idea",id:"Task 0",type:"task",progress:100,isDisabled:!0,styles:{progressColor:"#ffbb54",progressSelectedColor:"#ff9e0d"}},{start:new Date(2020,1,3),end:new Date(2020,1,5),name:"Idea123",id:"Task 1",type:"task",progress:100,isDisabled:!0,styles:{progressColor:"#ffbb54",progressSelectedColor:"#ff9e0d"}},{start:new Date(2020,1,1),end:new Date(2020,1,2),name:"Idea868",id:"Task 2",type:"task",progress:100,isDisabled:!0,styles:{progressColor:"#ffbb54",progressSelectedColor:"#ff9e0d"}}];function d(){return(0,i.jsx)(a.A,{tasks:l})}var x=s(8201),c=s(9021),p=s(5121),h=s(6242),g=s(657),j=s(8557),u=s(2280),f=s(8895),m=s(2797),b=s(6571),y=e=>{let{teamid:t}=e;console.log(t);let[s,o]=(0,r.useState)([{name:"extension 1",desc:"description 1",status:"enabled"},{name:"extension 2",desc:"description",status:"disabled"}]);return(0,i.jsx)("div",{style:{width:"80%",display:"flex",gap:"10px",rowGap:"10px",padding:"10px"},children:s.map(e=>(0,i.jsxs)(h.Z,{children:[(0,i.jsx)("h1",{children:e.name}),(0,i.jsx)("h5",{children:e.desc}),(0,i.jsx)(n.Z,{children:"View Details"}),(0,i.jsx)(n.Z,{children:"enabled"===e.status?"Disable":"Enable"})]}))})};let k=async e=>{try{await p.Z.post("".concat("http://localhost:8001","/start/").concat(e,"/").concat(8080+e))}catch(e){throw console.error("Error starting team:",e),e}},v=async e=>{try{await p.Z.post("".concat("http://localhost:8001","/stop/").concat(e))}catch(e){throw console.error("Error stopping team:",e),e}},F=async e=>{try{let t=await p.Z.get("".concat("http://localhost:8001","/logs/").concat(e));return t.data}catch(e){throw console.error("Error fetching logs:",e),e}};var w=e=>{let{teamID:t}=e,[s,n]=(0,r.useState)([]);return(0,r.useEffect)(()=>{let e=async()=>{try{let{logs:e}=await F(t),s=e.split("\x01");n(s)}catch(e){console.error("Error fetching logs:",e)}};e()},[t]),(0,i.jsx)(h.Z,{style:{width:"90%",alignSelf:"center",padding:"10px",backgroundColor:"#ddd",fontFamily:'"Courier New", monospace'},children:0===s.length?"Loading logs...":s.map((e,t)=>(0,i.jsx)("p",{children:e},t))})};function D(e){let[t,s]=(0,r.useState)(-1),[a,l]=(0,r.useState)(!1),[d,x]=(0,r.useState)(!1),[c,F]=(0,r.useState)([]),[D,C]=(0,r.useState)([]);(0,r.useEffect)(()=>{s(new URLSearchParams(window.location.search).get("id"))},[]),(0,r.useEffect)(()=>{(async()=>{console.log("Getting all servers");try{let e=(await p.Z.get("".concat("http://localhost:8001","/servers"))).data;t in Object.keys(e)&&l(!0)}catch(e){console.error("Could not reach server controller:",e)}})()},[t]);let S=async()=>{try{void 0!==t&&(a?await v(parseInt(t)):await k(parseInt(t)),l(!a))}catch(e){console.error("Error:",e)}};return(0,i.jsx)("div",{style:{width:"100vw",minHeight:"100vh",backgroundColor:"#f6f6f6"},children:(0,i.jsx)(h.Z,{sx:{display:"flex",flexDirection:"column",flexWrap:"nowrap",alignItems:"start",gap:"10px",width:"80%",margin:"auto",padding:2},children:-1!==t&&(0,i.jsxs)(i.Fragment,{children:[(0,i.jsxs)("span",{children:[(0,i.jsx)(n.Z,{variant:"contained",style:{display:"inline"},onClick:()=>{window.history.back()},children:"Back"}),(0,i.jsx)("h1",{style:{display:"inline",marginLeft:"30px"},children:"Team ".concat(t," server")})]}),(0,i.jsx)(n.Z,{variant:"contained",onClick:()=>{x(!0)},children:"Configure installed extensions"}),(0,i.jsx)(g.Z,{open:d,onClose:()=>{x(!d)},fullWidth:!0,maxWidth:"lg",children:(0,i.jsxs)(o.Z,{sx:{width:"100%"},children:[(0,i.jsxs)(j.Z,{children:["Team ",t," server extensions"]}),void 0!==t?(0,i.jsx)(y,{teamid:parseInt(t)}):null]})}),(0,i.jsxs)(u.Z,{sx:{width:"100%"},children:[(0,i.jsx)(f.Z,{expandIcon:(0,i.jsx)(b.Z,{}),children:(0,i.jsx)("h5",{children:"Advanced Configuration"})}),(0,i.jsx)(m.Z,{children:(0,i.jsxs)(o.Z,{sx:{display:"flex",flexDirection:"column",alignItems:"start",gap:"10px",width:"100%"},children:[(0,i.jsxs)("h3",{children:["URL to server: ","google.com"]}),(0,i.jsx)("h3",{children:"Logs: "}),-1!==t&&void 0!==t&&a?(0,i.jsx)(w,{teamID:parseInt(t)}):null,(0,i.jsx)(n.Z,{variant:"contained",onClick:S,children:a?"Force Stop":"Start"})]})})]})]})})})}let C=e=>{let{milestone:t,isExpanded:s,onClick:n}=e,[a,l]=(0,r.useState)(!1),[d,c]=(0,r.useState)(""),[p,h]=(0,r.useState)(""),g={position:"absolute",left:"40px",top:"125%",transform:"translateY(-20%)",width:"40px",height:"40px"},j={position:"absolute",left:"20px",top:"130%",transform:"translateY(-20%)",width:"40px",height:"40px",color:"red"},u={position:"absolute",left:"0%",top:"125%",transform:"translateY(-20%)",width:"40px",height:"40px",color:"blue"};return(0,i.jsxs)(o.Z,{border:"1px solid #ddd",borderRadius:"10px",p:2,sx:{backgroundColor:"#FFFFFF",marginTop:"16px",maxWidth:"100%",boxShadow:"0px 0px 5px 0px #D1D1D1",cursor:"pointer"},children:[(0,i.jsx)("div",{style:{display:"flex",flexDirection:"column",alignItems:"stretch"},children:(0,i.jsx)("div",{style:{marginBottom:"16px",display:"flex",flexDirection:"column"},onClick:n,children:(0,i.jsxs)("div",{style:{display:"flex",flexDirection:"row",justifyContent:"space-between",alignItems:"center"},children:[(0,i.jsxs)("div",{style:{display:"flex",flexDirection:"column",alignItems:"flex-start",flex:1},children:[(0,i.jsx)("div",{children:(0,i.jsx)("h2",{children:t.title})}),(0,i.jsxs)("div",{style:{display:"flex",flexDirection:"row",alignItems:"center"},children:[(0,i.jsxs)("span",{style:{fontSize:"20px",color:"#555",marginBottom:"8px"},children:[(0,i.jsx)("b",{children:"Project Description:"})," ",t.project.description]}),(0,i.jsxs)("span",{style:{fontSize:"20px",color:"#555",marginBottom:"8px",marginLeft:"16px"},children:[(0,i.jsx)("b",{children:"Expected Completion:"})," ",t.project.expectedCompletion]})]})]}),(0,i.jsxs)("div",{style:{display:"flex",alignItems:"center"},children:[(0,i.jsx)("button",{onClick:()=>{l(!0)},style:{fontSize:"16px",padding:"8px",backgroundColor:"#1977d2",color:"white",borderRadius:"5px",cursor:"pointer"},children:"Add Task"}),(0,i.jsx)("span",{style:{transform:s?"rotate(90deg)":"rotate(0deg)",fontSize:"24px",transition:"transform 0.3s ease",marginLeft:"8px"},children:"▶"})]})]})})}),s&&(0,i.jsx)("div",{style:{marginTop:"8px"},children:t.tasks.map(e=>(0,i.jsxs)("div",{style:{display:"flex",alignItems:"center",backgroundColor:"#FFFFFF",boxShadow:"0px 0px 5px 0px #D1D1D1",padding:"16px",borderRadius:"10px",marginTop:"16px"},children:[(0,i.jsxs)("div",{style:{flex:1},children:[(0,i.jsx)("h2",{children:e.title}),(0,i.jsx)("p",{children:e.description}),(0,i.jsxs)("div",{style:{position:"relative"},children:[(0,i.jsx)(x.Z,{style:g}),(0,i.jsx)(x.Z,{style:j}),(0,i.jsx)(x.Z,{style:u})]}),(0,i.jsx)("p",{children:(0,i.jsx)("br",{})})]}),(0,i.jsx)("p",{style:{color:"Orange",fontSize:"30px"},children:"Pending"})]},"task-".concat(e.id)))}),(0,i.jsx)("div",{style:{display:"flex",flexDirection:"column",alignItems:"stretch"},children:(0,i.jsx)("div",{style:{marginBottom:"16px",display:"flex",flexDirection:"column"},onClick:n})}),a&&(0,i.jsxs)("div",{style:{position:"fixed",top:"50%",left:"50%",transform:"translate(-50%, -50%)",boxShadow:"inherit",backgroundColor:"white",padding:"16px",borderRadius:"10px",zIndex:999},children:[(0,i.jsx)("h2",{children:"Add Task"}),(0,i.jsx)("label",{htmlFor:"taskName",children:"Task Name:"}),(0,i.jsx)("input",{type:"text",id:"taskName",value:d,onChange:e=>c(e.target.value)}),(0,i.jsx)("br",{}),(0,i.jsx)("br",{}),(0,i.jsx)("label",{htmlFor:"taskDescription",children:"Task Description:"}),(0,i.jsx)("textarea",{id:"taskDescription",value:p,onChange:e=>h(e.target.value)}),(0,i.jsx)("br",{}),(0,i.jsx)("br",{}),(0,i.jsx)("button",{onClick:()=>{c(""),h(""),l(!1)},children:"Add Task"}),(0,i.jsx)("button",{onClick:()=>l(!1),children:"Cancel"})]})]})};var S=()=>{let[e,t]=(0,r.useState)(!1),[s,a]=(0,r.useState)(!1),[l,x]=(0,r.useState)([]),[p,h]=(0,r.useState)(!1),[g,j]=(0,r.useState)("Project Management"),u=(e,t)=>{j(t)},f=e=>{let t=[...l];t[e]=!t[e],x(t)};return(0,i.jsxs)("div",{style:{display:"flex",flexDirection:"column"},children:[(0,i.jsx)(c.Z,{}),(0,i.jsxs)("div",{style:{backgroundColor:"#f6f6f6",width:"100%",minHeight:"100vh",padding:"16px",overflowY:"auto"},children:[(0,i.jsxs)("div",{style:{height:"100px",width:"100%"},children:[(0,i.jsx)(n.Z,{variant:"Project Management"===g?"contained":"outlined",onClick:e=>u(e,"Project Management"),style:{marginRight:"8px",width:"150px",borderRadius:"40px"},children:"Project Management"}),(0,i.jsx)(n.Z,{variant:"Server Management"===g?"contained":"outlined",onClick:e=>u(e,"Server Management"),style:{width:"150px",borderRadius:"40px"},children:"Server Management"})]}),(0,i.jsxs)("div",{style:{display:"flex"},children:["Project Management"===g&&(0,i.jsx)(i.Fragment,{children:(0,i.jsxs)("div",{style:{display:"flex",overflowY:"auto"},children:[(0,i.jsxs)("div",{style:{flex:"0 0 63%",maxWidth:"63%",marginRight:"16px",marginTop:"16px",backgroundColor:"#FBF8F8",boxShadow:"0px 0px 5px 0px #D1D1D1",padding:"16px",borderRadius:"10px",overflowY:"auto"},children:[(0,i.jsx)("div",{style:{borderRadius:"10px",backgroundColor:"#FFFFFF",width:"100%",boxShadow:"0px 0px 5px 0px #D1D1D1"},children:(0,i.jsx)(d,{style:{width:"100%"}})}),(0,i.jsxs)(o.Z,{border:"1px solid #ddd",borderRadius:"10px",p:2,sx:{backgroundColor:"#FFFFFF",marginTop:"16px",maxWidth:"100%",boxShadow:"0px 0px 5px 0px #D1D1D1",cursor:"pointer",overflowY:"auto"},children:[(0,i.jsxs)("div",{style:{display:"flex",justifyContent:"space-between",alignItems:"center"},onClick:()=>{h(!p)},children:[(0,i.jsx)("h2",{children:"Milestones"}),(0,i.jsx)("span",{style:{transform:p?"rotate(90deg)":"rotate(0deg)",transition:"transform 0.3s ease"},children:"▶"})]}),p&&(0,i.jsx)("div",{style:{marginTop:"8px"},children:[{id:1,title:"Milestone 1",project:{description:"Project 1 description...",expectedCompletion:"2 days"},tasks:[{id:1,title:"Task 1",description:"Task 1 description..."},{id:2,title:"Task 2",description:"Task 2 description..."},{id:3,title:"Task 3",description:"Task 3 description..."}]},{id:2,title:"Milestone 2",project:{description:"Project 2 description...",expectedCompletion:"3 days"},tasks:[{id:1,title:"Task 1",description:"Task 1 description..."},{id:2,title:"Task 2",description:"Task 2 description..."},{id:3,title:"Task 3",description:"Task 3 description..."}]},{id:1,title:"Milestone 3",project:{description:"Project 3 description...",expectedCompletion:"4 days"},tasks:[{id:1,title:"Task 1",description:"Task 1 description..."},{id:2,title:"Task 2",description:"Task 2 description..."},{id:3,title:"Task 3",description:"Task 3 description..."}]}].map((e,t)=>(0,i.jsx)(C,{milestone:e,isExpanded:l[t],onClick:()=>f(t)},"milestoneLeftCard".concat(e.id)))})]})]}),(0,i.jsxs)("div",{style:{flex:"0 0 32%",maxWidth:"32%",marginRight:"32px",marginLeft:"auto",backgroundColor:"#FBF8F8",boxShadow:"0px 0px 5px 0px #D1D1D1",padding:"16px",borderRadius:"10px",overflowY:"auto"},children:[(0,i.jsxs)(o.Z,{border:"1px solid #ddd",borderRadius:"10px",p:2,sx:{backgroundColor:"#FFFFFF",maxWidth:"100%",boxShadow:"0px 0px 5px 0px #D1D1D1"},children:[(0,i.jsx)("h2",{children:"Next Milestone Deadline"}),(0,i.jsxs)("p",{style:{fontSize:"20px"},children:["Name: ",(0,i.jsx)("span",{style:{fontSize:"14px"},children:"Project 1"})]}),(0,i.jsxs)("p",{style:{fontSize:"20px"},children:["Deadline: ",(0,i.jsx)("span",{style:{fontSize:"14px"},children:"2021-10-20"})]}),(0,i.jsxs)("p",{style:{fontSize:"20px"},children:["Time: ",(0,i.jsx)("span",{style:{fontSize:"14px"},children:"2d 3h 5m"})]})]}),(0,i.jsxs)(o.Z,{border:"1px solid #ddd",borderRadius:"10px",p:2,sx:{backgroundColor:"#FFFFFF",maxWidth:"100%",boxShadow:"0px 0px 5px 0px #D1D1D1"},children:[(0,i.jsx)("h2",{children:"Recived Payments"}),(0,i.jsx)("p",{style:{fontSize:"40px",textAlign:"center"},children:" $2000"})]}),(0,i.jsxs)(o.Z,{border:"1px solid #ddd",borderRadius:"10px",p:2,sx:{backgroundColor:"#FFFFFF",marginTop:"16px",maxWidth:"100%",boxShadow:"0px 0px 5px 0px #D1D1D1",cursor:"pointer"},onClick:()=>{t(!e)},children:[(0,i.jsxs)("div",{style:{display:"flex",justifyContent:"space-between",alignItems:"center"},children:[(0,i.jsx)("h2",{children:"Active Tasks"}),(0,i.jsx)("span",{style:{transform:e?"rotate(90deg)":"rotate(0deg)",transition:"transform 0.3s ease"},children:"▶"})]}),e&&(0,i.jsxs)("div",{style:{marginTop:"8px"},children:[(0,i.jsx)("div",{style:{backgroundColor:"#FFFFFF",boxShadow:"0px 0px 5px 0px #D1D1D1",padding:"16px",borderRadius:"10px"},children:(0,i.jsx)("h2",{children:"Task 1"})}),(0,i.jsx)("div",{style:{backgroundColor:"#FFFFFF",boxShadow:"0px 0px 5px 0px #D1D1D1",padding:"16px",borderRadius:"10px",marginTop:"16px"},children:(0,i.jsx)("h2",{children:"Task 2"})}),(0,i.jsx)("div",{style:{backgroundColor:"#FFFFFF",boxShadow:"0px 0px 5px 0px #D1D1D1",padding:"16px",borderRadius:"10px",marginTop:"16px"},children:(0,i.jsx)("h2",{children:"Task 3"})})]})]})]})]})}),"Server Management"===g&&(0,i.jsx)(i.Fragment,{children:(0,i.jsx)("div",{style:{display:"flex",width:"100%"},children:(0,i.jsx)(D,{})})})]})]})]})}}},function(e){e.O(0,[752,30,999,774,888,179],function(){return e(e.s=6413)}),_N_E=e.O()}]);
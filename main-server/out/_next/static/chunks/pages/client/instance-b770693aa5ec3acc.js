(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[276],{6413:function(e,t,n){(window.__NEXT_P=window.__NEXT_P||[]).push(["/client/instance",function(){return n(2739)}])},2739:function(e,t,n){"use strict";n.r(t),n.d(t,{default:function(){return y}});var s=n(5893),a=n(6242),r=n(9417),i=n(657),l=n(5616),c=n(8557),o=n(2280),d=n(8895),h=n(2797),x=n(6571),p=n(7294),u=e=>{let{teamid:t}=e;console.log(t);let[n,i]=(0,p.useState)([{name:"extension 1",desc:"description 1",status:"enabled"},{name:"extension 2",desc:"description",status:"disabled"}]);return(0,s.jsx)("div",{style:{width:"80%",display:"flex",gap:"10px",rowGap:"10px",padding:"10px"},children:n.map(e=>(0,s.jsxs)(a.Z,{children:[(0,s.jsx)("h1",{children:e.name}),(0,s.jsx)("h5",{children:e.desc}),(0,s.jsx)(r.Z,{children:"View Details"}),(0,s.jsx)(r.Z,{children:"enabled"===e.status?"Disable":"Enable"})]}))})},g=n(5121);let f=async e=>{try{await g.Z.post("http://localhost:8000/start/".concat(e,"/").concat(8080+e))}catch(e){throw console.error("Error starting team:",e),e}},w=async e=>{try{await g.Z.post("http://localhost:8000/stop/".concat(e))}catch(e){throw console.error("Error stopping team:",e),e}},j=async e=>{try{let t=await g.Z.get("http://localhost:8000/logs/".concat(e));return t.data}catch(e){throw console.error("Error fetching logs:",e),e}};var m=e=>{let{teamID:t}=e,[n,r]=(0,p.useState)([]);return(0,p.useEffect)(()=>{let e=async()=>{try{let{logs:e}=await j(t),n=e.split("\x01");r(n)}catch(e){console.error("Error fetching logs:",e)}};e()},[t]),(0,s.jsx)(a.Z,{style:{width:"90%",alignSelf:"center",padding:"10px",backgroundColor:"#ddd",fontFamily:'"Courier New", monospace'},children:0===n.length?"Loading logs...":n.map((e,t)=>(0,s.jsx)("p",{children:e},t))})};function y(e){let[t,n]=(0,p.useState)(0),[g,j]=(0,p.useState)(()=>{let e=localStorage.getItem("team".concat(t,"_started"));return!!e&&JSON.parse(e)}),[y,Z]=(0,p.useState)(!1),[v,S]=(0,p.useState)([]),[E,_]=(0,p.useState)([]);(0,p.useEffect)(()=>{n(new URLSearchParams(window.location.search).get("id"))},[]),(0,p.useEffect)(()=>{localStorage.setItem("team".concat(t,"_started"),JSON.stringify(g))},[g,t]);let k=async()=>{try{void 0!==t&&(g?await w(parseInt(t)):await f(parseInt(t)),j(!g))}catch(e){console.error("Error:",e)}};return(0,s.jsx)("div",{style:{width:"100vw",minHeight:"100vh",backgroundColor:"#f6f6f6"},children:(0,s.jsxs)(a.Z,{sx:{display:"flex",flexDirection:"column",alignItems:"start",gap:"10px",width:"80%",height:"80%",margin:"auto",padding:2},children:[(0,s.jsxs)("span",{children:[(0,s.jsx)(r.Z,{variant:"contained",style:{display:"inline"},onClick:()=>{window.history.back()},children:"Back"}),(0,s.jsx)("h1",{style:{display:"inline",marginLeft:"30px"},children:"Team ".concat(t," server")})]}),(0,s.jsx)(r.Z,{variant:"contained",onClick:()=>{Z(!0)},children:"Configure installed extensions"}),(0,s.jsx)(i.Z,{open:y,onClose:()=>{Z(!y)},fullWidth:!0,maxWidth:"lg",children:(0,s.jsxs)(l.Z,{sx:{width:"100%"},children:[(0,s.jsxs)(c.Z,{children:["Team ",t," server extensions"]}),void 0!==t?(0,s.jsx)(u,{teamid:parseInt(t)}):null]})}),(0,s.jsx)(l.Z,{sx:{width:"100%"},children:(0,s.jsx)("h2",{children:"Team Files"})}),(0,s.jsxs)(o.Z,{sx:{width:"100%"},children:[(0,s.jsx)(d.Z,{expandIcon:(0,s.jsx)(x.Z,{}),children:(0,s.jsx)("h5",{children:"Advanced Configuration"})}),(0,s.jsx)(h.Z,{children:(0,s.jsxs)(l.Z,{sx:{display:"flex",flexDirection:"column",alignItems:"start",gap:"10px",width:"100%"},children:[(0,s.jsxs)("h3",{children:["URL to server: ","google.com"]}),(0,s.jsx)("h3",{children:"Logs: "}),void 0!==t?(0,s.jsx)(m,{teamID:parseInt(t)}):null,(0,s.jsx)(r.Z,{variant:"contained",onClick:k,children:g?"Force Stop":"Start"})]})})]})]})})}}},function(e){e.O(0,[193,283,947,511,774,888,179],function(){return e(e.s=6413)}),_N_E=e.O()}]);
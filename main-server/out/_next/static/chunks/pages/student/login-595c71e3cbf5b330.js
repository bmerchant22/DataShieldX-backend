(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[258],{3914:function(e,t,n){(window.__NEXT_P=window.__NEXT_P||[]).push(["/student/login",function(){return n(1963)}])},1963:function(e,t,n){"use strict";n.r(t),n.d(t,{default:function(){return y}});var o=n(5893),r=n(7294),a=n(9417),s=n(7462),i=n(1657),l=n(7297);let c=(e,t)=>(0,s.Z)({WebkitFontSmoothing:"antialiased",MozOsxFontSmoothing:"grayscale",boxSizing:"border-box",WebkitTextSizeAdjust:"100%"},t&&!e.vars&&{colorScheme:e.palette.mode}),d=e=>(0,s.Z)({color:(e.vars||e).palette.text.primary},e.typography.body1,{backgroundColor:(e.vars||e).palette.background.default,"@media print":{backgroundColor:(e.vars||e).palette.common.white}}),u=(e,t=!1)=>{var n;let o={};t&&e.colorSchemes&&Object.entries(e.colorSchemes).forEach(([t,n])=>{var r;o[e.getColorSchemeSelector(t).replace(/\s*&/,"")]={colorScheme:null==(r=n.palette)?void 0:r.mode}});let r=(0,s.Z)({html:c(e,t),"*, *::before, *::after":{boxSizing:"inherit"},"strong, b":{fontWeight:e.typography.fontWeightBold},body:(0,s.Z)({margin:0},d(e),{"&::backdrop":{backgroundColor:(e.vars||e).palette.background.default}})},o),a=null==(n=e.components)||null==(n=n.MuiCssBaseline)?void 0:n.styleOverrides;return a&&(r=[r,a]),r};var m=function(e){let t=(0,i.Z)({props:e,name:"MuiCssBaseline"}),{children:n,enableColorScheme:a=!1}=t;return(0,o.jsxs)(r.Fragment,{children:[(0,o.jsx)(l.Z,{styles:e=>u(e,a)}),n]})},p=n(2914),h=n(629),g=n(5616),b=n(6886),x=n(5861),f=n(3156);function y(){let[e,t]=(0,r.useState)(!1),n=async e=>{e.preventDefault();let n=new FormData(e.currentTarget);try{let e=await fetch("".concat("http://localhost:8000","/student/auth"),{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({team:parseInt(n.get("team")),username:n.get("username"),password:n.get("password")})});e.ok?(t(!1),window.location.assign("/student/dashboard")):t(!0)}catch(e){console.error("Error during login:",e),t(!0)}};return(0,o.jsx)(f.Z,{component:"main",maxWidth:"lg",children:(0,o.jsx)(g.Z,{sx:{marginTop:8},children:(0,o.jsxs)(b.ZP,{container:!0,children:[(0,o.jsx)(m,{}),(0,o.jsx)(b.ZP,{item:!0,xs:12,sm:8,md:5,component:h.Z,elevation:6,square:!0,children:(0,o.jsxs)(g.Z,{sx:{my:8,mx:1,display:"flex",flexDirection:"column",alignItems:"center"},children:[(0,o.jsx)(x.Z,{component:"h1",variant:"h5",children:"Sign in"}),(0,o.jsxs)("form",{noValidate:!0,onSubmit:n,style:{height:"100%"},children:[(0,o.jsx)(p.Z,{margin:"normal",required:!0,fullWidth:!0,id:"team",label:"Team",name:"team",autoComplete:"team",autoFocus:!0}),(0,o.jsx)(p.Z,{margin:"normal",required:!0,fullWidth:!0,id:"username",label:"Username",name:"username",autoComplete:"username",autoFocus:!0}),(0,o.jsx)(p.Z,{margin:"normal",required:!0,fullWidth:!0,name:"password",label:"Password",type:"password",id:"password",autoComplete:"current-password"}),(0,o.jsx)(a.Z,{type:"submit",variant:"contained",sx:{mt:3,mb:2},children:"Sign In"})]})]})}),(0,o.jsx)(b.ZP,{item:!0,xs:!1,sm:4,md:7,sx:{backgroundImage:"url(https://source.unsplash.com/random)",backgroundRepeat:"no-repeat",backgroundColor:e=>"light"===e.palette.mode?e.palette.grey[50]:e.palette.grey[900],backgroundSize:"cover",backgroundPosition:"center"}})]})})})}}},function(e){e.O(0,[193,886,283,199,774,888,179],function(){return e(e.s=3914)}),_N_E=e.O()}]);
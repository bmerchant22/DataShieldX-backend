(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[258],{8241:function(e,t,o){(window.__NEXT_P=window.__NEXT_P||[]).push(["/student/login",function(){return o(9145)}])},9145:function(e,t,o){"use strict";o.r(t),o.d(t,{default:function(){return y}});var n=o(4246),r=o(7378),a=o(353),s=o(5773),i=o(8014),l=o(4840);let c=(e,t)=>(0,s.Z)({WebkitFontSmoothing:"antialiased",MozOsxFontSmoothing:"grayscale",boxSizing:"border-box",WebkitTextSizeAdjust:"100%"},t&&!e.vars&&{colorScheme:e.palette.mode}),u=e=>(0,s.Z)({color:(e.vars||e).palette.text.primary},e.typography.body1,{backgroundColor:(e.vars||e).palette.background.default,"@media print":{backgroundColor:(e.vars||e).palette.common.white}}),d=(e,t=!1)=>{var o;let n={};t&&e.colorSchemes&&Object.entries(e.colorSchemes).forEach(([t,o])=>{var r;n[e.getColorSchemeSelector(t).replace(/\s*&/,"")]={colorScheme:null==(r=o.palette)?void 0:r.mode}});let r=(0,s.Z)({html:c(e,t),"*, *::before, *::after":{boxSizing:"inherit"},"strong, b":{fontWeight:e.typography.fontWeightBold},body:(0,s.Z)({margin:0},u(e),{"&::backdrop":{backgroundColor:(e.vars||e).palette.background.default}})},n),a=null==(o=e.components)||null==(o=o.MuiCssBaseline)?void 0:o.styleOverrides;return a&&(r=[r,a]),r};var m=function(e){let t=(0,i.Z)({props:e,name:"MuiCssBaseline"}),{children:o,enableColorScheme:a=!1}=t;return(0,n.jsxs)(r.Fragment,{children:[(0,n.jsx)(l.Z,{styles:e=>d(e,a)}),o]})},p=o(6428),h=o(920),g=o(1703),b=o(8544),x=o(2750),f=o(309);function y(){let[e,t]=(0,r.useState)(!1),o=async e=>{e.preventDefault();let o=new FormData(e.currentTarget);try{let e=await fetch("".concat("http://localhost:8001","/api/auth"),{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({team:o.get("team"),user:o.get("username"),pass:o.get("password")})});e.ok?(t(!1),window.location.assign("/student/dashboard/")):t(!0)}catch(e){console.error("Error during login:",e),t(!0)}};return(0,n.jsx)(f.Z,{component:"main",maxWidth:"lg",children:(0,n.jsx)(g.Z,{sx:{marginTop:8},children:(0,n.jsxs)(b.ZP,{container:!0,children:[(0,n.jsx)(m,{}),(0,n.jsx)(b.ZP,{item:!0,xs:12,sm:8,md:5,component:h.Z,elevation:6,square:!0,children:(0,n.jsxs)(g.Z,{sx:{my:8,mx:1,display:"flex",flexDirection:"column",alignItems:"center"},children:[(0,n.jsx)(x.Z,{component:"h1",variant:"h5",children:"Sign in"}),(0,n.jsxs)("form",{noValidate:!0,onSubmit:o,style:{height:"100%"},children:[(0,n.jsx)(p.Z,{margin:"normal",required:!0,fullWidth:!0,id:"team",label:"Team",name:"team",autoComplete:"team",autoFocus:!0}),(0,n.jsx)(p.Z,{margin:"normal",required:!0,fullWidth:!0,id:"username",label:"Username",name:"username",autoComplete:"username",autoFocus:!0}),(0,n.jsx)(p.Z,{margin:"normal",required:!0,fullWidth:!0,name:"password",label:"Password",type:"password",id:"password",autoComplete:"current-password"}),(0,n.jsx)(a.Z,{type:"submit",variant:"contained",sx:{mt:3,mb:2},children:"Sign In"})]})]})}),(0,n.jsx)(b.ZP,{item:!0,xs:!1,sm:4,md:7,sx:{backgroundImage:"url(https://source.unsplash.com/random)",backgroundRepeat:"no-repeat",backgroundColor:e=>"light"===e.palette.mode?e.palette.grey[50]:e.palette.grey[900],backgroundSize:"cover",backgroundPosition:"center"}})]})})})}}},function(e){e.O(0,[338,830,774,888,179],function(){return e(e.s=8241)}),_N_E=e.O()}]);
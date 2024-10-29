"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[1014],{8024:(e,n,o)=>{o.r(n),o.d(n,{assets:()=>c,contentTitle:()=>s,default:()=>p,frontMatter:()=>r,metadata:()=>a,toc:()=>d});var t=o(5893),i=o(1151);const r={sidebar_position:4,title:"Offboarding Checklist"},s="Consumer Offboarding",a={id:"consumer-development/offboarding",title:"Offboarding Checklist",description:"To offboard a consumer chain simply submit a ConsumerRemovalProposal governance proposal listing a stop_time. After stop time passes, the provider chain will remove the chain from the ICS protocol (it will stop sending validator set updates).",source:"@site/versioned_docs/version-v4.5.0/consumer-development/offboarding.md",sourceDirName:"consumer-development",slug:"/consumer-development/offboarding",permalink:"/interchain-security/v4.5.0/consumer-development/offboarding",draft:!1,unlisted:!1,tags:[],version:"v4.5.0",sidebarPosition:4,frontMatter:{sidebar_position:4,title:"Offboarding Checklist"},sidebar:"tutorialSidebar",previous:{title:"Onboarding Checklist",permalink:"/interchain-security/v4.5.0/consumer-development/onboarding"},next:{title:"Changeover Procedure",permalink:"/interchain-security/v4.5.0/consumer-development/changeover-procedure"}},c={},d=[];function l(e){const n={code:"code",h1:"h1",p:"p",pre:"pre",...(0,i.a)(),...e.components};return(0,t.jsxs)(t.Fragment,{children:[(0,t.jsx)(n.h1,{id:"consumer-offboarding",children:"Consumer Offboarding"}),"\n",(0,t.jsxs)(n.p,{children:["To offboard a consumer chain simply submit a ",(0,t.jsx)(n.code,{children:"ConsumerRemovalProposal"})," governance proposal listing a ",(0,t.jsx)(n.code,{children:"stop_time"}),". After stop time passes, the provider chain will remove the chain from the ICS protocol (it will stop sending validator set updates)."]}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-js",children:'// ConsumerRemovalProposal is a governance proposal on the provider chain to remove (and stop) a consumer chain.\n// If it passes, all the consumer chain\'s state is removed from the provider chain. The outstanding unbonding\n// operation funds are released.\n{\n    // the title of the proposal\n    "title": "This was a great chain",\n    "description": "Here is a .md formatted string specifying removal details",\n    // the chain-id of the consumer chain to be stopped\n    "chain_id": "consumerchain-1",\n    // the time on the provider chain at which all validators are responsible to stop their consumer chain validator node\n    "stop_time": "2023-03-07T12:40:00.000000Z",\n}\n'})}),"\n",(0,t.jsx)(n.p,{children:"More information will be listed in a future version of this document."})]})}function p(e={}){const{wrapper:n}={...(0,i.a)(),...e.components};return n?(0,t.jsx)(n,{...e,children:(0,t.jsx)(l,{...e})}):l(e)}},1151:(e,n,o)=>{o.d(n,{Z:()=>a,a:()=>s});var t=o(7294);const i={},r=t.createContext(i);function s(e){const n=t.useContext(r);return t.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function a(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(i):e.components||i:s(e.components),t.createElement(r.Provider,{value:n},e.children)}}}]);
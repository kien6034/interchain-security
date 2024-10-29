"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[6627],{2444:(e,n,r)=>{r.r(n),r.d(n,{assets:()=>d,contentTitle:()=>o,default:()=>l,frontMatter:()=>s,metadata:()=>a,toc:()=>c});var i=r(5893),t=r(1151);const s={sidebar_position:2},o="Reward Distribution",a={id:"features/reward-distribution",title:"Reward Distribution",description:"Sending and distributing rewards from consumer chains to the provider chain is handled by the Reward Distribution sub-protocol.",source:"@site/versioned_docs/version-v4.5.0/features/reward-distribution.md",sourceDirName:"features",slug:"/features/reward-distribution",permalink:"/interchain-security/v4.5.0/features/reward-distribution",draft:!1,unlisted:!1,tags:[],version:"v4.5.0",sidebarPosition:2,frontMatter:{sidebar_position:2},sidebar:"tutorialSidebar",previous:{title:"Key Assignment",permalink:"/interchain-security/v4.5.0/features/key-assignment"},next:{title:"ICS Provider Proposals",permalink:"/interchain-security/v4.5.0/features/proposals"}},d={},c=[{value:"Whitelisting Reward Denoms",id:"whitelisting-reward-denoms",level:2}];function h(e){const n={a:"a",admonition:"admonition",code:"code",h1:"h1",h2:"h2",p:"p",pre:"pre",...(0,t.a)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(n.h1,{id:"reward-distribution",children:"Reward Distribution"}),"\n",(0,i.jsxs)(n.p,{children:["Sending and distributing rewards from consumer chains to the provider chain is handled by the ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc/blob/main/spec/app/ics-028-cross-chain-validation/overview_and_basic_concepts.md#reward-distribution",children:"Reward Distribution sub-protocol"}),"."]}),"\n",(0,i.jsxs)(n.p,{children:["Consumer chains have the option of sharing (a portion of) their block rewards (inflation tokens and fees) with the provider chain validators and delegators.\nIn Interchain Security, block rewards are periodically sent from the consumer to the provider according to consumer chain parameters using an IBC transfer channel.\nThis channel is created during consumer chain initialization, unless it is provided via the ",(0,i.jsx)(n.code,{children:"ConsumerAdditionProposal"})," when adding a new consumer chain.\nFor more details, see the ",(0,i.jsx)(n.a,{href:"/interchain-security/v4.5.0/introduction/params#reward-distribution-parameters",children:"reward distribution parameters"}),"."]}),"\n",(0,i.jsx)(n.admonition,{type:"tip",children:(0,i.jsxs)(n.p,{children:["Providing an IBC transfer channel (see ",(0,i.jsx)(n.code,{children:"DistributionTransmissionChannel"}),") enables a consumer chain to re-use one of the existing channels to the provider for consumer chain rewards distribution. This will preserve the ",(0,i.jsx)(n.code,{children:"ibc denom"})," that may already be in use.\nThis is especially important for standalone chains transitioning to become consumer chains.\nFor more details, see the ",(0,i.jsx)(n.a,{href:"/interchain-security/v4.5.0/consumer-development/changeover-procedure",children:"changeover procedure"}),"."]})}),"\n",(0,i.jsx)(n.p,{children:"Reward distribution on the provider is handled by the distribution module."}),"\n",(0,i.jsx)(n.h2,{id:"whitelisting-reward-denoms",children:"Whitelisting Reward Denoms"}),"\n",(0,i.jsxs)(n.p,{children:["The ICS distribution system works by allowing consumer chains to send rewards to a module address on the provider called the ",(0,i.jsx)(n.code,{children:"ConsumerRewardsPool"}),".\nTo avoid spam, the provider must whitelist denoms before accepting them as ICS rewards.\nOnly whitelisted denoms are transferred from the ",(0,i.jsx)(n.code,{children:"ConsumerRewardsPool"})," to the ",(0,i.jsx)(n.code,{children:"FeePoolAddress"}),", to be distributed to delegators and validators.\nThe whitelisted denoms can be adjusted through governance by sending a ",(0,i.jsx)(n.a,{href:"/interchain-security/v4.5.0/features/proposals#changerewarddenomproposal",children:(0,i.jsx)(n.code,{children:"ChangeRewardDenomProposal"})}),"."]}),"\n",(0,i.jsx)(n.p,{children:"To query the list of whitelisted reward denoms on the Cosmos Hub, use the following command:"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-bash",children:"> gaiad q provider registered-consumer-reward-denoms\ndenoms:\n- ibc/0025F8A87464A471E66B234C4F93AEC5B4DA3D42D7986451A059273426290DD5\n- ibc/6B8A3F5C2AD51CD6171FA41A7E8C35AD594AB69226438DB94450436EA57B3A89\n- uatom\n"})}),"\n",(0,i.jsxs)(n.admonition,{type:"tip",children:[(0,i.jsxs)(n.p,{children:["Use the following command to get a human readable denom from the ",(0,i.jsx)(n.code,{children:"ibc/*"})," denom trace format:"]}),(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-bash",children:">  gaiad query ibc-transfer denom-trace ibc/0025F8A87464A471E66B234C4F93AEC5B4DA3D42D7986451A059273426290DD5\ndenom_trace:\n  base_denom: untrn\n  path: transfer/channel-569\n"})})]})]})}function l(e={}){const{wrapper:n}={...(0,t.a)(),...e.components};return n?(0,i.jsx)(n,{...e,children:(0,i.jsx)(h,{...e})}):h(e)}},1151:(e,n,r)=>{r.d(n,{Z:()=>a,a:()=>o});var i=r(7294);const t={},s=i.createContext(t);function o(e){const n=i.useContext(s);return i.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function a(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(t):e.components||t:o(e.components),i.createElement(s.Provider,{value:n},e.children)}}}]);
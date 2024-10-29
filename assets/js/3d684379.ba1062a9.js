"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[3597],{8409:(e,i,n)=>{n.r(i),n.d(i,{assets:()=>c,contentTitle:()=>t,default:()=>h,frontMatter:()=>o,metadata:()=>d,toc:()=>a});var r=n(5893),s=n(1151);const o={sidebar_position:3},t="Interchain Security Parameters",d={id:"introduction/params",title:"Interchain Security Parameters",description:"The parameters necessary for Interchain Security (ICS) are defined in",source:"@site/versioned_docs/version-v4.5.0/introduction/params.md",sourceDirName:"introduction",slug:"/introduction/params",permalink:"/interchain-security/v4.5.0/introduction/params",draft:!1,unlisted:!1,tags:[],version:"v4.5.0",sidebarPosition:3,frontMatter:{sidebar_position:3},sidebar:"tutorialSidebar",previous:{title:"Terminology",permalink:"/interchain-security/v4.5.0/introduction/terminology"},next:{title:"Technical Specification",permalink:"/interchain-security/v4.5.0/introduction/technical-specification"}},c={},a=[{value:"Time-Based Parameters",id:"time-based-parameters",level:2},{value:"ProviderUnbondingPeriod",id:"providerunbondingperiod",level:3},{value:"ConsumerUnbondingPeriod",id:"consumerunbondingperiod",level:3},{value:"TrustingPeriodFraction",id:"trustingperiodfraction",level:3},{value:"CCVTimeoutPeriod",id:"ccvtimeoutperiod",level:3},{value:"InitTimeoutPeriod",id:"inittimeoutperiod",level:3},{value:"VscTimeoutPeriod",id:"vsctimeoutperiod",level:3},{value:"BlocksPerDistributionTransmission",id:"blocksperdistributiontransmission",level:3},{value:"TransferPeriodTimeout",id:"transferperiodtimeout",level:3},{value:"Reward Distribution Parameters",id:"reward-distribution-parameters",level:2},{value:"ConsumerRedistributionFraction",id:"consumerredistributionfraction",level:3},{value:"BlocksPerDistributionTransmission",id:"blocksperdistributiontransmission-1",level:3},{value:"TransferTimeoutPeriod",id:"transfertimeoutperiod",level:3},{value:"DistributionTransmissionChannel",id:"distributiontransmissionchannel",level:3},{value:"ProviderFeePoolAddrStr",id:"providerfeepooladdrstr",level:3},{value:"Slash Throttle Parameters",id:"slash-throttle-parameters",level:2},{value:"SlashMeterReplenishPeriod",id:"slashmeterreplenishperiod",level:3},{value:"SlashMeterReplenishFraction",id:"slashmeterreplenishfraction",level:3},{value:"MaxThrottledPackets",id:"maxthrottledpackets",level:3},{value:"RetryDelayPeriod",id:"retrydelayperiod",level:3},{value:"Epoch Parameters",id:"epoch-parameters",level:2},{value:"BlocksPerEpoch",id:"blocksperepoch",level:3}];function l(e){const i={a:"a",admonition:"admonition",code:"code",h1:"h1",h2:"h2",h3:"h3",li:"li",p:"p",pre:"pre",strong:"strong",ul:"ul",...(0,s.a)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(i.h1,{id:"interchain-security-parameters",children:"Interchain Security Parameters"}),"\n",(0,r.jsx)(i.p,{children:"The parameters necessary for Interchain Security (ICS) are defined in"}),"\n",(0,r.jsxs)(i.ul,{children:["\n",(0,r.jsxs)(i.li,{children:["the ",(0,r.jsx)(i.code,{children:"Params"})," structure in ",(0,r.jsx)(i.code,{children:"proto/interchain_security/ccv/provider/v1/provider.proto"})," for the provider;"]}),"\n",(0,r.jsxs)(i.li,{children:["the ",(0,r.jsx)(i.code,{children:"Params"})," structure in ",(0,r.jsx)(i.code,{children:"proto/interchain_security/ccv/consumer/v1/consumer.proto"})," for the consumer."]}),"\n"]}),"\n",(0,r.jsx)(i.h2,{id:"time-based-parameters",children:"Time-Based Parameters"}),"\n",(0,r.jsx)(i.p,{children:"ICS relies on the following time-based parameters."}),"\n",(0,r.jsx)(i.h3,{id:"providerunbondingperiod",children:"ProviderUnbondingPeriod"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"ProviderUnbondingPeriod"})," is the unbonding period on the provider chain as configured during chain genesis. This parameter can later be changed via governance."]}),"\n",(0,r.jsx)(i.h3,{id:"consumerunbondingperiod",children:"ConsumerUnbondingPeriod"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"ConsumerUnbondingPeriod"})," is the unbonding period on the consumer chain."]}),"\n",(0,r.jsxs)(i.admonition,{type:"info",children:[(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"ConsumerUnbondingPeriod"})," is set via the ",(0,r.jsx)(i.code,{children:"ConsumerAdditionProposal"})," governance proposal to add a new consumer chain.\nIt is recommended that every consumer chain set and unbonding period shorter than ",(0,r.jsx)(i.code,{children:"ProviderUnbondingPeriod"})]}),(0,r.jsx)("br",{}),(0,r.jsx)(i.p,{children:"Example:"}),(0,r.jsx)(i.pre,{children:(0,r.jsx)(i.code,{children:"ConsumerUnbondingPeriod = ProviderUnbondingPeriod - one day\n"})})]}),"\n",(0,r.jsx)(i.p,{children:"Unbonding operations (such as undelegations) are completed on the provider only after the unbonding period elapses on every consumer."}),"\n",(0,r.jsx)(i.h3,{id:"trustingperiodfraction",children:"TrustingPeriodFraction"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"TrustingPeriodFraction"})," is used to calculate the ",(0,r.jsx)(i.code,{children:"TrustingPeriod"})," of created IBC clients on both provider and consumer chains."]}),"\n",(0,r.jsxs)(i.p,{children:["Setting ",(0,r.jsx)(i.code,{children:"TrustingPeriodFraction"})," to ",(0,r.jsx)(i.code,{children:"0.5"})," would result in the following:"]}),"\n",(0,r.jsx)(i.pre,{children:(0,r.jsx)(i.code,{children:"TrustingPeriodFraction = 0.5\nProviderClientOnConsumerTrustingPeriod = ProviderUnbondingPeriod * 0.5\nConsumerClientOnProviderTrustingPeriod = ConsumerUnbondingPeriod * 0.5\n"})}),"\n",(0,r.jsxs)(i.p,{children:["Note that a light clients must be updated within the ",(0,r.jsx)(i.code,{children:"TrustingPeriod"})," in order to avoid being frozen."]}),"\n",(0,r.jsxs)(i.p,{children:["For more details, see the ",(0,r.jsx)(i.a,{href:"https://github.com/cosmos/ibc/blob/main/spec/client/ics-007-tendermint-client/README.md",children:"IBC specification of Tendermint clients"}),"."]}),"\n",(0,r.jsx)(i.h3,{id:"ccvtimeoutperiod",children:"CCVTimeoutPeriod"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"CCVTimeoutPeriod"})," is the period used to compute the timeout timestamp when sending IBC packets."]}),"\n",(0,r.jsxs)(i.p,{children:["For more details, see the ",(0,r.jsx)(i.a,{href:"https://github.com/cosmos/ibc/blob/main/spec/core/ics-004-channel-and-packet-semantics/README.md#sending-packets",children:"IBC specification of Channel & Packet Semantics"}),"."]}),"\n",(0,r.jsx)(i.admonition,{type:"warning",children:(0,r.jsx)(i.p,{children:"If a sent packet is not relayed within this period, then the packet times out. The CCV channel used by the interchain security protocol is closed, and the corresponding consumer is removed."})}),"\n",(0,r.jsx)(i.p,{children:"CCVTimeoutPeriod may have different values on the provider and consumer chains."}),"\n",(0,r.jsxs)(i.ul,{children:["\n",(0,r.jsxs)(i.li,{children:[(0,r.jsx)(i.code,{children:"CCVTimeoutPeriod"})," on the provider ",(0,r.jsx)(i.strong,{children:"must"})," be larger than ",(0,r.jsx)(i.code,{children:"ConsumerUnbondingPeriod"})]}),"\n",(0,r.jsxs)(i.li,{children:[(0,r.jsx)(i.code,{children:"CCVTimeoutPeriod"})," on the consumer is initial set via the ",(0,r.jsx)(i.code,{children:"ConsumerAdditionProposal"})]}),"\n"]}),"\n",(0,r.jsx)(i.h3,{id:"inittimeoutperiod",children:"InitTimeoutPeriod"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"InitTimeoutPeriod"})," is the maximum allowed duration for CCV channel initialization to execute."]}),"\n",(0,r.jsxs)(i.p,{children:["For any consumer chain, if the CCV channel is not established within ",(0,r.jsx)(i.code,{children:"InitTimeoutPeriod"})," then the consumer chain will be removed and therefore will not be secured by the provider chain."]}),"\n",(0,r.jsxs)(i.p,{children:["The countdown starts when the ",(0,r.jsx)(i.code,{children:"spawn_time"})," specified in the ",(0,r.jsx)(i.code,{children:"ConsumerAdditionProposal"})," is reached."]}),"\n",(0,r.jsx)(i.h3,{id:"vsctimeoutperiod",children:"VscTimeoutPeriod"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"VscTimeoutPeriod"})," is the provider-side param that enables the provider to timeout VSC packets even when a consumer chain is not live.\nIf the ",(0,r.jsx)(i.code,{children:"VscTimeoutPeriod"})," is ever reached for a consumer chain that chain will be considered not live and removed from interchain security."]}),"\n",(0,r.jsx)(i.admonition,{type:"tip",children:(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"VscTimeoutPeriod"})," MUST be larger than the ",(0,r.jsx)(i.code,{children:"ConsumerUnbondingPeriod"}),"."]})}),"\n",(0,r.jsx)(i.h3,{id:"blocksperdistributiontransmission",children:"BlocksPerDistributionTransmission"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"BlocksPerDistributionTransmission"})," is the number of blocks between rewards transfers from the consumer to the provider."]}),"\n",(0,r.jsx)(i.h3,{id:"transferperiodtimeout",children:"TransferPeriodTimeout"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"TransferPeriodTimeout"})," is the period used to compute the timeout timestamp when sending IBC transfer packets from a consumer to the provider."]}),"\n",(0,r.jsxs)(i.p,{children:["If this timeout expires, then the transfer is attempted again after ",(0,r.jsx)(i.code,{children:"BlocksPerDistributionTransmission"})," blocks."]}),"\n",(0,r.jsxs)(i.ul,{children:["\n",(0,r.jsxs)(i.li,{children:[(0,r.jsx)(i.code,{children:"TransferPeriodTimeout"})," on the consumer is initial set via the ",(0,r.jsx)(i.code,{children:"ConsumerAdditionProposal"})," gov proposal to add the consumer"]}),"\n",(0,r.jsxs)(i.li,{children:[(0,r.jsx)(i.code,{children:"TransferPeriodTimeout"})," should be smaller than ",(0,r.jsx)(i.code,{children:"BlocksPerDistributionTransmission x avg_block_time"})]}),"\n"]}),"\n",(0,r.jsx)(i.h2,{id:"reward-distribution-parameters",children:"Reward Distribution Parameters"}),"\n",(0,r.jsx)(i.admonition,{type:"tip",children:(0,r.jsxs)(i.p,{children:["The following chain parameters dictate consumer chain distribution amount and frequency.\nThey are set at consumer genesis and ",(0,r.jsx)(i.code,{children:"BlocksPerDistributionTransmission"}),", ",(0,r.jsx)(i.code,{children:"ConsumerRedistributionFraction"}),"\n",(0,r.jsx)(i.code,{children:"TransferTimeoutPeriod"})," must be provided in every ",(0,r.jsx)(i.code,{children:"ConsumerChainAddition"})," proposal."]})}),"\n",(0,r.jsx)(i.h3,{id:"consumerredistributionfraction",children:"ConsumerRedistributionFraction"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"ConsumerRedistributionFraction"})," is the fraction of tokens allocated to the consumer redistribution address during distribution events. The fraction is a string representing a decimal number. For example ",(0,r.jsx)(i.code,{children:'"0.75"'})," would represent ",(0,r.jsx)(i.code,{children:"75%"}),"."]}),"\n",(0,r.jsxs)(i.admonition,{type:"tip",children:[(0,r.jsx)(i.p,{children:"Example:"}),(0,r.jsxs)(i.p,{children:["With ",(0,r.jsx)(i.code,{children:"ConsumerRedistributionFraction"})," set to ",(0,r.jsx)(i.code,{children:'"0.75"'})," the consumer chain would send ",(0,r.jsx)(i.code,{children:"75%"})," of its block rewards and accumulated fees to the consumer redistribution address, and the remaining ",(0,r.jsx)(i.code,{children:"25%"})," to the provider chain every ",(0,r.jsx)(i.code,{children:"BlocksPerDistributionTransmission"})," blocks."]})]}),"\n",(0,r.jsx)(i.h3,{id:"blocksperdistributiontransmission-1",children:"BlocksPerDistributionTransmission"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"BlocksPerDistributionTransmission"})," is the number of blocks between IBC token transfers from the consumer chain to the provider chain."]}),"\n",(0,r.jsx)(i.h3,{id:"transfertimeoutperiod",children:"TransferTimeoutPeriod"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"TransferTimeoutPeriod"})," is the timeout period for consumer chain reward distribution IBC packets."]}),"\n",(0,r.jsx)(i.h3,{id:"distributiontransmissionchannel",children:"DistributionTransmissionChannel"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"DistributionTransmissionChannel"})," is the provider chain IBC channel used for receiving consumer chain reward distribution token transfers. This is automatically set during the consumer-provider handshake procedure."]}),"\n",(0,r.jsx)(i.h3,{id:"providerfeepooladdrstr",children:"ProviderFeePoolAddrStr"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"ProviderFeePoolAddrStr"})," is the provider chain fee pool address used for receiving consumer chain reward distribution token transfers. This is automatically set during the consumer-provider handshake procedure."]}),"\n",(0,r.jsx)(i.h2,{id:"slash-throttle-parameters",children:"Slash Throttle Parameters"}),"\n",(0,r.jsx)(i.h3,{id:"slashmeterreplenishperiod",children:"SlashMeterReplenishPeriod"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"SlashMeterReplenishPeriod"})," exists on the provider such that once the slash meter becomes not-full, the slash meter is replenished after this period has elapsed."]}),"\n",(0,r.jsxs)(i.p,{children:["The meter is replenished to an amount equal to the slash meter allowance for that block, or ",(0,r.jsx)(i.code,{children:"SlashMeterReplenishFraction * CurrentTotalVotingPower"}),"."]}),"\n",(0,r.jsx)(i.h3,{id:"slashmeterreplenishfraction",children:"SlashMeterReplenishFraction"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"SlashMeterReplenishFraction"})," exists on the provider as the portion (in range [0, 1]) of total voting power that is replenished to the slash meter when a replenishment occurs."]}),"\n",(0,r.jsxs)(i.p,{children:["This param also serves as a maximum fraction of total voting power that the slash meter can hold. The param is set/persisted as a string, and converted to a ",(0,r.jsx)(i.code,{children:"sdk.Dec"})," when used."]}),"\n",(0,r.jsx)(i.h3,{id:"maxthrottledpackets",children:"MaxThrottledPackets"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"MaxThrottledPackets"})," exists on the provider as the maximum amount of throttled slash or vsc matured packets that can be queued from a single consumer before the provider chain halts, it should be set to a large value."]}),"\n",(0,r.jsx)(i.p,{children:"This param would allow provider binaries to panic deterministically in the event that packet throttling results in a large amount of state-bloat. In such a scenario, packet throttling could prevent a violation of safety caused by a malicious consumer, at the cost of provider liveness."}),"\n",(0,r.jsx)(i.admonition,{type:"info",children:(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"MaxThrottledPackets"})," was deprecated in ICS versions >= v3.2.0 due to the implementation of ",(0,r.jsx)(i.a,{href:"/interchain-security/v4.5.0/adrs/adr-008-throttle-retries",children:"ADR-008"}),"."]})}),"\n",(0,r.jsx)(i.h3,{id:"retrydelayperiod",children:"RetryDelayPeriod"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"RetryDelayPeriod"})," exists on the consumer for ",(0,r.jsx)(i.strong,{children:"ICS versions >= v3.2.0"})," (introduced by the implementation of ",(0,r.jsx)(i.a,{href:"/interchain-security/v4.5.0/adrs/adr-008-throttle-retries",children:"ADR-008"}),") and is the period at which the consumer retries to send a ",(0,r.jsx)(i.code,{children:"SlashPacket"})," that was rejected by the provider."]}),"\n",(0,r.jsx)(i.h2,{id:"epoch-parameters",children:"Epoch Parameters"}),"\n",(0,r.jsx)(i.h3,{id:"blocksperepoch",children:"BlocksPerEpoch"}),"\n",(0,r.jsxs)(i.p,{children:[(0,r.jsx)(i.code,{children:"BlocksPerEpoch"})," exists on the provider for ",(0,r.jsx)(i.strong,{children:"ICS versions >= 3.3.0"})," (introduced by the implementation of ",(0,r.jsx)(i.a,{href:"/interchain-security/v4.5.0/adrs/adr-014-epochs",children:"ADR-014"}),")\nand corresponds to the number of blocks that constitute an epoch. This param is set to 600 by default. Assuming we need 6 seconds to\ncommit a block, the duration of an epoch corresponds to 1 hour. This means that a ",(0,r.jsx)(i.code,{children:"VSCPacket"})," would be sent to a consumer\nchain once at the end of every epoch, so once every 600 blocks. This parameter can be adjusted via a governance proposal,\nhowever careful consideration is needed so that ",(0,r.jsx)(i.code,{children:"BlocksPerEpoch"})," is not too large. A large ",(0,r.jsx)(i.code,{children:"BlocksPerEpoch"})," could lead to a delay\nof ",(0,r.jsx)(i.code,{children:"VSCPacket"}),"s and hence potentially lead to ",(0,r.jsx)(i.a,{href:"https://informal.systems/blog/learning-to-live-with-unbonding-pausing",children:"unbonding pausing"}),".\nFor setting ",(0,r.jsx)(i.code,{children:"BlocksPerEpoch"}),", we also need to consider potential slow chain upgrades that could delay the sending of a\n",(0,r.jsx)(i.code,{children:"VSCPacket"}),", as well as potential increases in the time it takes to commit a block (e.g., from 6 seconds to 30 seconds)."]})]})}function h(e={}){const{wrapper:i}={...(0,s.a)(),...e.components};return i?(0,r.jsx)(i,{...e,children:(0,r.jsx)(l,{...e})}):l(e)}},1151:(e,i,n)=>{n.d(i,{Z:()=>d,a:()=>t});var r=n(7294);const s={},o=r.createContext(s);function t(e){const i=r.useContext(o);return r.useMemo((function(){return"function"==typeof e?e(i):{...i,...e}}),[i,e])}function d(e){let i;return i=e.disableParentContext?"function"==typeof e.components?e.components(s):e.components||s:t(e.components),r.createElement(o.Provider,{value:i},e.children)}}}]);

# Hyperledger Fabric

Hyperledger Fabric 是 Hyperledger 项目的一个组成部分，是一个区块链框架的实现。它将成为区块链应用开发、解决方案的基础。Fabric 框架支持组件化、可插拔的共识服务（Consensus Service）、成员服务（Membership Service）；"许可（Permissioned）"特性使之为"私有性"、"保密性"提供了可靠的解决方案；智能合约（Smart Contracts）在 Fabric 中通过"Chaincode"得以实现。

Hyperledger Composer 也是 Hyperledger 项目的一个组成部分。通过它，人们可以更快速、容易地建立区块链业务模型，进行区块链网络及应用的开发、部署，并与现有系统、数据进行集成。

联盟链、溯源、


-------






















-------

 

从本质上来说，Hyperledger可以分为三类：分布式账本技术（基本上就是区块链）、库以及工具。接下来就按上述顺序来介绍，认真看哦~


# 分布式账本技术


DLT

Besu算是最新的工具，来自ConsenSys，被称为Pantheon。它十分有趣，这是一个基于Java的以太坊客户端，实现了以太坊虚拟机（EVM ，即Ethereum Virtual Machine），该虚拟机支持许可网络以及公共网络，包括Ropsten、Rinkeby和Görli等测试网络。

Besu支持的共识算法包括工作量证明（PoW）、权威证明（PoA）和IBFT2，IBFT2是个基于p2p的协议。Besu包括web套接字、HTTP和命令行界面，用于以太坊网络的工作和交互，并且支持智能合约。

Besu地址：
https://www.hyperledger.org/projects/besu

Burrow是一个模块化区块链客户端，其中包含一个经过许可的智能合约解释器，该解释器的部分开发采用了以太坊虚拟机（EVM）规范，旨在运行EVM智能合约。它利用了Tendermint证明共识引擎。

Tendermint地址：https://tendermint.com/

该项目作出了高交易吞吐量的承诺，因此各组织可以基于此构建EVM智能合约并在本地进行部署。坦白来讲，这个项目在过去的一年里似乎并没有取得很大的进展，github上的大多数更新似乎都是无关紧要的，我也没怎么仔细研究。

Burrow地址：
https://www.hyperledger.org/projects/hyperledger-burrow

Fabric可能算是Hyperledger体系中最重要的项目。它十分强大，处于运行状态，且具备非常多功能。Fabric是一个经过许可的企业级DLT框架，采用模块化设计，在使用时具有很强的灵活性。我自己曾在游戏市场中使用过它，也确实有被惊艳到。

基本上来说，你可以在你自己的计算机上部署自己的区块链，并控制成本和环境。

Fabric地址：
https://www.hyperledger.org/projects/fabric

Indy可是相当的有趣，它在分布式账本的基础上提供了一个自我主权身份生态系统，其试图建立一个可跨其他软件系统使用的区块链存储身份系统。我很喜欢这个想法，而且它似乎也正在积极开发中。

Indy地址：
https://www.hyperledger.org/projects/hyperledger-indy

Iroha也很特别。一年前我曾在一个游戏项目中考虑过它，它本可以很完美，但在当时那个时候，还远远不够完美。它旨在通过一小组快速命令和查询来操纵账户和数字资产。

验证节点可以使用Gossip协议来分发半签名交易，作为多重签名交易的一部分。当分类账本状态存储在PostgreSQL中时，区块存储在文件中。如果用发散性思维去想的话，这个项目会有无限的可能性。

Iroha地址：
https://www.hyperledger.org/projects/iroha

Sawtooth是去年我为我的游戏平台考虑过的另一个技术项目，但当时，它并不是我们心中想做的合适的模型，这也是我们最终选择了Fabric的原因。它最初使用的是“所用时间证明Proof of Elapsed Time”共识算法，并且它需要一些运行在计算机上的Intel软件来进行管理，但现在你有了诸如RAFT和PBFT等多种选择。

Sawtooth真正的好处是可以使用多种语言来编写智能合约，这让我想起了Dragonchain。Sawtooth可以通过SETH（Sawtooth/Ethereum）来执行以太坊智能合约。如果你正在考虑部署经许可的区块链的话，那么Fabric和Sawtooth是值得一看的。


# 库

Aries是基于区块链点对点交互基础设施。它并不是一条区块链，也不是一个应用程序，同时也没有投入使用。其目标是为不同的去中心化系统提供点对点交互、加密管理、可验证的信息交换和安全消息传递服务。

它和Hyperledger Indy项目以及Ursa项目都有联系。其中有一些想法非常有趣，但目前还没有投入运行。

Aries地址：
https://www.hyperledger.org/projects/aries

Hyperledger Indy项目：
https://www.hyperledger.org/projects/hyperledger-indy

Ursa项目：
https://www.hyperledger.org/projects/ursa

Quilt提供了使用跨账本协议（Interledger Protocol）在不同账本系统间进行相互操作的能力，该协议通常被用在支付场景。通过提供甚至支持非区块链系统的原子交换，该库被用于在账本系统间进行价值传递。

看看Git，该项目似乎正在被积极的研究，对于需要这项功能的人来说可是件好事。

Quilt地址：
https://www.hyperledger.org/projects/quilt

Transact是个非常新的项目，甚至都没有针对它的Git项目。其理念是要创造一个用于执行智能合约的标准接口，该接口是从实际的分布式账本实现中抽象出来的。

我喜欢这个想法的原因是它会简化整个智能合约过程，并将其开放给其他语言使用。虽然目前不太确定它进行到哪一步了，但是它背后的家伙却是意志十分坚定的。这也是个你在开发中要常考虑的项目。

Transact地址：
https://www.hyperledger.org/projects/transact

Ursa是一个共享的加密库，旨在避免重复的加密工作（让人们使用同一个库），同时也以提高安全性为目标。它是使用C和Rust构建的，并且有着减少冗余工作的光荣使命。从Git中很难看出其开发和广泛使用的情况，但是离发布1.0版本还很遥远，而且自2019年4月以来就没有发布过新版本。

Ursa地址：
https://www.hyperledger.org/projects/ursa


# 工具

Caliper是一个性能测量工具，用于衡量预定义用例中特定区块链实施的性能。它会生成具有多项性能指标的报告，例如TPS（每秒交易量Transactions Per Second）、交易延迟、资源利用率等等。

这是个十分有用的工具，可以让你了解技术的实施情况以及可能需要改进的地方。据Git显示，它在很多Hyperledger项目中都是十分有效的。

Caliper地址：
https://www.hyperledger.org/projects/caliper

Cello将按需的“即服务”的部署模型引入了区块链，从而使创建、管理和停止区块链都变得更加容易。它在各种基础设施（如裸机、虚拟机和容器）上提供多租户链服务。

这有点像Docker Swarm或Kubernetes，只不过它是用在区块链上的。该项目经常和Hyperledger Explorer资源管理器结合使用。

Hyperledger Explorer资源管理器：
https://www.hyperledger.org/projects/explorer

截至2019年8月，鉴于Fabric 1.4+的原因，Composer项目（https://www.hyperledger.org/projects/composer）已被弃用。虽然这让很多事情处理起来更加方便，但我还是挺吃惊的，毕竟去年我经常使用它。

Cello地址：
https://www.hyperledger.org/projects/cello

Composer项目：
https://www.hyperledger.org/projects/composer

Explorer是一个用户友好的Web应用程序，它可以查看、调用、部署或查询区块、交易和相关数据、网络信息、链码和交易族（transaction families）以及储存在分类账本中的其他相关信息。如果你在Hyperledger世界中工作的话，这是一个能和你成为朋友的APP，十分受大家的欢迎。

Explorer地址：
https://www.hyperledger.org/projects/explorer

最后是Grid，这是构建供应链解决方案的框架。我不打算深入地探讨它，这是一个框架和库的集合，以此共同构建供应链。如果你正在关注供应链的话，那么不妨看一看，但目前它还没有准备好迎接黄金阶段。

Grid地址：
https://grid.hyperledger.org/

看完本文，你应该了解了一些非常有趣的项目，其中一些是实时可用的，而有些则不是。我真的很喜欢Hyperledger宇宙中正在发生的事情，而且这绝对值得一看。
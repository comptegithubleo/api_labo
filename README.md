## Laboratoire Cyber
Cyberlab pour Proxmox, espaces virtuels inter-connectés pour collaborer sur des projets, déployé avec Docker Compose et scripts bash.
Chaque utilisateur à son propre espace fermé (on premise ou depuis un VPN), et peut lancer et travailler sur ses projets.
Les utilisateurs peuvent s'inviter entre eux, et collaborer ensemble (inter-VLAN)
- Un utilisateur "professeur" peut inviter plusieurs utilisateurs "étudiants" et travailler sur un objectif de cours
- Un utilisateur "CTF Admin" peut inviter plusieurs utilisateurs "CTF Players"
- Plusieurs utilisateurs "élèves" peuvent s'inviter entre eux pour créer un réseau maillé et créer des groupes de travails

### Structure
- Chaque utilisateur à un compte (identifiant **integer** + accès Wireguard + compte Proxmox (allocation ressources / templates, etc))
- Chaque création de compte génère un sous-réseau lié à l'identifiant (id 12 -> 10.0.12.0/24) et initialise le routage réseau pour celui-ci
- Un utilisateur peut accéder au Proxmox GUI et au webserver "client"
- L'administrateur peut accéder au Proxmox GUI et au webserver "admin"

### Webserver Client :
- Un utilisateur peut voir tous les utilisateurs existants sur le serveur, ainsi que ceux connectés à son espace
- Un utilisateur peut supprimer les connexions existantes avec d'autres utilisateurs et fermer son espace

### Webserver Admin :
- Un administrateur peut ajouter, modifier, ou supprimer des liens entre utilisateurs

### Router :
- Le router comprend une API appelée par les webservers client et admin, et appelle à son tour les APIs Proxmox et WGDashboard pour la création de comptes, modifications de liens, etc
- Les ajouts, modifications, suppressions de liens entre utilisateurs se font par nftables, en routant les pairs Wireguard IP avec leur VLAN Proxmox respectifs.

...


# Groupe
David Chau, Daniel Monteiro, Brian MADAPATHAGE SENANAYAKE, Prashath Sivayanama

# Flight Aggregator — Résumé
Le projet **Flight Aggregator** permet d’agréger des vols provenant de deux APIs, puis de les trier selon différents critères.

---

## Fonctionnalités principales

- Serveur Go (`http.Server`) avec deux routes :
  - **GET `/health`** → vérifie la santé du serveur
  - **GET `/flights`** → retourne tous les vols agrégés
- Récupération des vols depuis deux APIs (`j-server1` et `j-server2`)
- Décodage JSON → Structs Go
- Deux repositories utilisant une interface commune
- Tri disponible :
  - `SortByPrice`
  - `SortByTimeTravel`
  - `SortByDepartureDate`
- Tests :
  - tests des algorithmes de tri
  - tests du service métier avec mocks (`testify`)

---

## Endpoints

### GET `/health`
Retourne `200 OK`.

### GET `/flights`
Retourne tous les vols agrégés.

### GET `/flights?filter=...`
Filtres disponibles :
- `/flights?filter=SortByPrice`
- `/flights?filter=SortByTimeTravel`
- `/flights?filter=SortByDepartureDate`

---

## Lancer le projet

```bash
docker compose build
docker compose up
```

## Lancer les tests

```bash
cd server && gotestsum
```

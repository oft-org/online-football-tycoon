# Tournaments and Seasons

This document describes the basic structure and logic for managing football competitions in the project.

## Entities

### 1. Tournament

Represents a competition that belongs to a specific country. It can be either a League or a Cup. Leagues can have multiple divisions with hierarchical promotion/relegation relationships.

**Fields:**
- `id`: Unique identifier.
- `name`: Name of the competition (e.g., "La Liga", "FA Cup").
- `type`: Either `League` or `Cup`.
- `country_code`: The country the tournament belongs to (ISO alpha-3).
- `division`: Division number (1 = top division).
- `promotion_to`: (Optional) Tournament ID to which teams are promoted.
- `descent_to`: (Optional) Tournament ID to which teams are relegated.

---

### 2. Season

Represents a specific edition of a tournament, such as "Premier League 2025/26".

**Fields:**
- `id`: Unique identifier.
- `tournament_id`: The tournament this season belongs to.
- `from_date`: Start date of the season.
- `to_date`: End date of the season.

---

### 3. SeasonTeam

Associates teams with a particular season of a tournament.

**Fields:**
- `season_id`: ID of the season.
- `team_id`: ID of the participating team.

---

## Functional Overview

### ➤ Creating a Tournament

Create a tournament with a name, type, and country. Leagues can include multiple divisions. You can link divisions using `promotion_to` and `descent_to`.

### ➤ Starting a New Season

Each tournament can have one or more seasons over time. A season defines the time range and teams that will participate.

### ➤ Assigning Teams to a Season

Teams are assigned to a season using the `season_team` table. This can be done:
- Automatically based on previous season results.
- Manually for testing or cup formats.

### ➤ Simulating Matches

You can create a separate `Match` entity linked to a season and teams. This would handle results, schedules, and stats.

### ➤ Ending a Season

At the end of a season:
- League standings determine promotions and relegations.
- New seasons can be initialized based on outcomes.

---

## Notes

- Promotion and relegation logic can be implemented based on league standings.
- Cup tournaments do not need divisions or promotion structure.
- Tournaments can exist independently per country.


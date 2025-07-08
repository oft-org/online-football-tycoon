# 📌 Note on Missing Pull Request Merge Commits (July 2025)

On July 8, 2025, a `git push --force` was applied to the `main` branch in order to replace its commit history with that of the `feature/competition-structure` branch.

This decision was made to resolve recurring merge issues and unify the codebase between both branches.

As a result:

- The code from all merged PRs (#22 to #26) is **fully present** and **functionally correct** in the `main` branch.
- However, the original **merge commits** associated with those PRs no longer appear in `main`'s commit history.
- The pull requests remain available on GitHub for reference and auditing purposes.

No code or feature was lost in the process — only the visibility of merge events in the commit log was affected.

If needed, the previous state of `main` can be recovered via local `git reflog` or backup tags.

---

✅ **Summary**  
- Code is intact  
- PRs #22 to #26 are merged  
- Merge commits are no longer visible in `main`  

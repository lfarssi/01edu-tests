## ✅ BASIC FUNCTIONAL TESTS

### 1. **Rename a file**

```bash
touch oldname.txt
mv oldname.txt newname.txt
ls newname.txt
```

✅ *Should show `newname.txt`, and `oldname.txt` should be gone.*

---

### 2. **Move file into directory**

```bash
mkdir mydir
touch file1.txt
mv file1.txt mydir/
ls mydir/
```

✅ *Should see `file1.txt` inside `mydir`*

---

### 3. **Move and rename file into a directory**

```bash
mv mydir/file1.txt mydir/renamed.txt
ls mydir/
```

✅ *Should see `renamed.txt`, and `file1.txt` should be gone*

---

### 4. **Rename a directory**

```bash
mkdir olddir
mv olddir newdir
ls -d newdir
```

✅ *Should see `newdir` and not `olddir`*

---

### 5. **Move directory into another directory**

```bash
mkdir parent
mv newdir parent/
ls parent/
```

✅ *Should see `newdir` inside `parent/`*

---

## ❌ ERROR TEST CASES (No flags)

### 6. **Move non-existent file**

```bash
mv not_a_file.txt dest.txt
```

❌ *Should show error: `No such file or directory`*

---

### 7. **Move into non-existent directory**

```bash
touch temp.txt
mv temp.txt nosuchdir/
```

❌ *Should show error: `No such file or directory`*

---

### 8. **Move file onto existing file (no overwrite handling)**

```bash
echo "one" > file_a.txt
echo "two" > file_b.txt
mv file_a.txt file_b.txt
cat file_b.txt
```

⚠️ *In a basic implementation, this would overwrite silently. You should test whether the content is `"one"` now. If not, overwriting didn’t work.*

---

### 9. **Missing arguments**

```bash
mv
mv fileonly
```

❌ *Should show: `missing operand` or `missing destination`*

---

## ⚠️ SPECIAL CASES

### 10. **Move multiple files (should fail in basic version)**

```bash
touch a.txt b.txt
mv a.txt b.txt newname.txt
```

❌ *Should fail: too many arguments*

---

### 11. **Moving file with same source and destination**

```bash
touch same.txt
mv same.txt same.txt
```

⚠️ *Should do nothing, or return: `file is the same`*

---

## 🔁 RENAMING & MOVING EDGE CASES

### 12. **Rename with spaces in the name**

```bash
touch "file with space.txt"
mv "file with space.txt" "renamed file.txt"
ls "renamed file.txt"
```

✅ *Should handle quotes properly and rename the file*

---

### 13. **Move to path ending with a slash but is not a dir**

```bash
touch file.txt
touch notadir
mv file.txt notadir/
```

❌ *Should fail: "not a directory" (unless `notadir` is a directory)*

---

### 14. **Move directory over a file**

```bash
mkdir dir1
touch file.txt
mv dir1 file.txt
```

❌ *Should fail: cannot overwrite file with directory*

---

### 15. **Move file over a directory**

```bash
touch file.txt
mkdir dir1
mv file.txt dir1
```

✅ *Should move file into `dir1/`*

---

### 16. **Move directory into itself (indirectly)**

```bash
mkdir -p a/b
mv a a/b/
```

❌ *Should fail or prevent loop: "cannot move directory into itself"*

---

### 17. **Move into existing directory with same file name**

```bash
mkdir dir1
touch dir1/file.txt
touch file.txt
mv file.txt dir1/
```

✅ *Should overwrite `dir1/file.txt` or keep both based on behavior (your shell likely overwrites)*

---

## 🧪 INVALID INPUTS & STRUCTURE

### 18. **Try to move nothing**

```bash
mv
```

❌ *Should return: "missing operand"*

---

### 19. **Try to move only one argument**

```bash
mv only_source.txt
```

❌ *Should return: "missing destination operand"*

---

### 20. **Trailing slashes in filenames (should still work)**

```bash
touch weirdfile
mv weirdfile weirdfile/
```

❌ *Should fail: `weirdfile/` is not a directory*

---

### 21. **Move file to a file path with a slash in middle (invalid path)**

```bash
touch a.txt
mv a.txt not/real/path.txt
```

❌ *Should fail: intermediate directory does not exist*

---

## ⚠️ SYSTEM-LIKE BEHAVIOR

### 22. **Move file with same name but different case (case-sensitive)**

```bash
touch lower.txt
mv lower.txt LOWER.txt
ls LOWER.txt
```

✅ *Should rename depending on case sensitivity of filesystem*

---

### 23. **Move a directory onto itself**

```bash
mkdir testdir
mv testdir testdir
```

❌ *Should do nothing or error: "invalid move"*

---

### 24. **Overwrite file with directory name**

```bash
touch mydir
mkdir actualdir
mv actualdir mydir
```

❌ *Should fail: cannot overwrite file with directory*

---

### 25. **Move files using relative paths**

```bash
mkdir parent
touch parent/file1
mv parent/file1 file1_moved
ls file1_moved
```

✅ *Should move properly from relative path*

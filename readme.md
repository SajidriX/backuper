Hello!

To use backuper, you need to compile it(and download at first).
=== Downloading guide ===
May be you do it yourself?

=== Compiling guide ===
1. Install golang.
sudo <package-manager> install golang
package manager:

Debian, Ubuntu, Mint, etc. --- apt
Fedora,RedHat, etc. --- dnf
Arch, Manjaro, Garuda --- pacman
2. Get to downloaded dir.

3.Compile.
go build main.go

You compiled it! Good!

======= USAGE =======
1. Run compiled file
2. Enter: backup <path_to_file_or_dir>
3.Wait

==== EXAMPLE ====
./main ~/.config/starship.toml
Backup created: /home/sajison/backups/starship.toml_2026-04-27_11-05-45


Now you know how to use it! Cool!

===== FEATURES =====
1.Can copy file and dir
2.Adds timestamp in format YYYY-MM-DD_HH-MM-SS to the backup name
3.Stores all backups in ~/backups

author - Sajison(sajidrix@gmail.com but I check my mail rarely, reddit - u/SajidriX)
license - MIT (check LICENSE.md)

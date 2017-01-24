if (Test-Path .\images) {
    $files = Get-ChildItem .\images | where {$_.Extension -eq ".jpg"}
    $random = Get-Random -Minimum 0 -Maximum ($files.Length)

    $abs = Resolve-Path (Join-Path "images\" $files[$random])
    .\wallpaper.exe $abs
}
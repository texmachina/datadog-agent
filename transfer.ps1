Param(
    [parameter(
    HelpMessage="Please pass in one of[agent, process-agent, system-probe, omnibus]")]
    [string]$name
)

$remote_host = "<YOUR_REMOTE_HOST>""

if ($name -eq "omnibus") {
    # Get the latest msi to be written to in the \omnibus\pkg directory
    $local_path = Get-ChildItem .\omnibus\pkg\*.msi | Sort-Object LastWriteTime | Select-Object -last 1
    $remote_address = $remote_host + ':C:\agent.msi'
} else {
    $local_path = '.\bin\' + $name + '\' + $name + '.exe'
    $remote_address = $remote_host + ':C:\' + $name + '.exe'
}

Write-Host "Copying " $local_path " to " $remote_address
scp.exe $local_path $remote_address
